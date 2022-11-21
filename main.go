package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//config  holds flags configurations
type config struct {
	ext  string //extension to filter out
	size int64  //min file size
	list bool   //list files
	del  bool   // delete file
}

func main() {
	root := flag.String("root", ".", "Root directory to start")
	//action options
	list := flag.Bool("list", false, "List files only")
	del := flag.Bool("del", false, "Delete files")

	//filter options
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	//create an instance with the flag values
	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
		del:  *del,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
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

			if cfg.del {
				return delFile(path)
			}

			// List is the default option if nothing else was set
			return listFile(path, out)
		})
}
