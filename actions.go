package main

import (
	"fmt"
	"io"
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

func delFile(path string) error {
	return os.Remove(path)
}
