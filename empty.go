package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	file     = flag.Bool("f", false, "Search file")
	absolute = flag.Bool("a", false, "Print absolute path")
)

func main() {
	os.Exit(search())
}

func search() int {
	flag.Parse()
	dirPath := "."

	if flag.NArg() > 0 {
		dirPath = filepath.Clean(flag.Arg(0))

		_, err := ioutil.ReadDir(dirPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
			return 1
		}
	}

	err := filepath.Walk(dirPath, func(path string, fileInfo os.FileInfo, err error) error {
		if *file {
			if fileInfo.IsDir() {
				return nil
			}

			if fileInfo.Size() != 0 {
				return nil
			}
		} else {
			if !fileInfo.IsDir() {
				return nil
			}

			files, err := ioutil.ReadDir(path)
			if err != nil {
				return err
			}

			if len(files) != 0 {
				return nil
			}
		}

		if *absolute {
			absPath, err := filepath.Abs(path)

			if err != nil {
				return err
			}
			path = absPath
		}

		fmt.Fprintln(os.Stdout, path)
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		return 1
	}
	return 0
}
