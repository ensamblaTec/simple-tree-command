package pkg

import (
	"fmt"
	"os"
)

// OpenPath create pointer

// ListDirectories list all directories in current path
func ListDirectories(path string) error {
	files, err := List(path)
	if err != nil {
		return err
	}
	index := 1
	for _, v := range files {
		if v.IsDir() {
			fmt.Printf("%d. %s\n", index, v.Name())
			index++
		}
	}
	return nil
}

// ListFiles list all files in current path
func ListFiles(path string) error {
	files, err := List(path)
	if err != nil {
		return err
	}
	index := 1
	for _, v := range files {
		if !v.IsDir() {
			fmt.Printf("%d. %s\n", index, v.Name())
			index++
		}
	}
	return nil
}

// Tree list all (file/directory) in current path and sub-folders
func Tree() {

}

// List all files in current directory
func List(path string) ([]os.DirEntry, error) {
	open, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return open, nil
}
