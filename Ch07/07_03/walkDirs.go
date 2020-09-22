// walk over the dir structure
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//absoulte path = pwd
	root, _ := filepath.Abs(".")
	fmt.Println("Processing path", root)
	//walk through the dir structure and list of dir and files
	err := filepath.Walk(root, processPath)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func processPath(path string, info os.FileInfo, err error) error {
	//check for walkfunc in fielpath
	if err != nil {
		return err
	}
	//check we are not working with root folder
	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
		}
	}
	return nil
}
