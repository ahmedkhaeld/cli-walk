package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

//filterOut checks if the current path has to be ignored or not
//requirements are file size must Not less than minSize, so it does not exclude,
//file ext  equal required ext
//filterOut return true to exclude a path that does not meet the requirements
//return false to NOT exclude a path that meets the requirements
//e.g. file size 12, minSize flag 13, so file is excluded from the listing
func filterOut(path, ext string, minSize int64, info os.FileInfo) bool {
	if info.IsDir() || info.Size() < minSize {
		return true
	}
	if ext != "" && filepath.Ext(path) != ext {
		return true
	}
	return false
}

//listFile prints out the path of the current file to out, then
//returning any potential error
func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}

//delFile takes in a path to be removed,
//l provide the destination of the logs
func delFile(path string, l *log.Logger) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	l.Println(path)
	return nil
}

//archiveFile takes in
//destDire: where the files will be archived;
//root: the root dir where search was started, use this value to create a similar tree in the destination
//;path: the path of the file to be archived
//;return a potential error
//
//archiveFile has two responsibilities:
//preserve the relative dir tree
//so the files are archived in the same directories relative to the source root
//,and to compress the data
//
//e.g.  $ go run . -root /tmp/gomisc/ -ext .go -archive /tmp/gomisc_bkp
//
//destDir : /tmp/go_misc_bkp; source /tmp/go_misc/
//
//targetPath: /tmp/go_misc_bkp/misc/reboot/reboot_test.go.gz
//
//combined from
//
//destDir: /tmp/go_misc_bkp/
//
// relDir: /misc/reboot/
//
// fileBase: reboot_test.go.
func archiveFile(destDir, root, path string) error {

	info, err := os.Stat(destDir)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", destDir)
	}

	//extract the relative directory tree of the file to be archived in
	//related to its source root path
	relDir, err := filepath.Rel(root, filepath.Dir(path))
	if err != nil {
		return err
	}
	//append to the fileBase  .gz extension
	fileBase := fmt.Sprintf("%s.gz", filepath.Base(path))

	//join all the three pieces together to generate target path
	targetPath := filepath.Join(destDir, relDir, fileBase)
	//e.g. targetPath /tmp/go_misc_bkp/misc/reboot/reboot_test.go.gz
	//destDir: /tmp/go_misc_bkp/
	// relDir: /misc/reboot/
	// fileBase: reboot_test.go.gz

	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return err
	}

	out, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	in, err := os.Open(path)
	if err != nil {
		return err
	}
	defer in.Close()

	zw := gzip.NewWriter(out)

	zw.Name = filepath.Base(path)

	if _, err = io.Copy(zw, in); err != nil {
		return err
	}

	if err := zw.Close(); err != nil {
		return err
	}

	return out.Close()
}
