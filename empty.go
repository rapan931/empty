package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	file = flag.Bool("f", false, "File only")
)

func main() {
	flag.Parse()
	searchPath := "."

	if flag.NArg() > 0 {
		searchPath = flag.Arg(0)

		_, err := ioutil.ReadDir(searchPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
			os.Exit(1)
		}
	}

	err := filepath.Walk(searchPath, func(path string, fileInfo os.FileInfo, err error) error {
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

		fmt.Println(path)
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
}
