package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

//config  holds flags configurations
type config struct {
	ext     string    //extension to filter out
	size    int64     //min file size
	list    bool      //list files
	del     bool      // delete file
	wLog    io.Writer //represent the log destination
	archive string    //archive dir
}

func main() {
	root := flag.String("root", ".", "Root directory to start")
	logFile := flag.String("log", "", "Log deletes to this file")
	//action options
	list := flag.Bool("list", false, "List files only")
	archive := flag.String("archive", "", "Archive directory")
	del := flag.Bool("del", false, "Delete files")

	//filter options
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	var (
		f   = os.Stdout
		err error
	)

	//create file that will store the logs
	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	//create an instance with the flag values
	c := config{
		ext:     *ext,
		size:    *size,
		list:    *list,
		del:     *del,
		wLog:    f, // inject filename
		archive: *archive,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	//delLogger define a logger with a file[io.Writer] as destination
	//as the file name provided from the cmd
	delLogger := log.New(cfg.wLog, "Deleted File: ", log.LstdFlags)

	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			//if filterOut true then ignore current path
			if filterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}

			// If list was explicitly set, don't do anything else
			if cfg.list {
				return listFile(path, out)
			}

			if cfg.archive != "" {
				if err := archiveFile(cfg.archive, root, path); err != nil {
					return err
				}
			}

			//inject the path to be deleted, and dest logs file
			if cfg.del {
				return delFile(path, delLogger)
			}

			// List is the default option if nothing else was set
			return listFile(path, out)
		})
}
