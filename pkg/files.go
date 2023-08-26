package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

// OpenPath create pointer

func ListAll(path string) error {
	files, err := List(path)
	if err != nil {
		return err
	}
	index := 1
	for _, v := range files {
		fmt.Printf("%d. %s\n", index, v.Name())
		index++
	}
	return nil
}

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
func Tree(path string, level int) error {
	files, err := List(path)
	if err != nil {
		return err
	}
	for _, v := range files {
		if v.IsDir() {
			indent := ""
			for i := 0; i < level; i++ {
				indent += "    "
			}
			fmt.Printf("%s[%s]\n", indent, v.Name())
			_ = Tree(filepath.Join(path, v.Name()), level+1)
		} else {
			indent := ""
			for i := 0; i < level; i++ {
				indent += "    "
			}
			fmt.Printf("%s%s\n", indent, v.Name())
		}
	}

	return nil
}

// List all files in current directory
func List(path string) ([]os.DirEntry, error) {
	open, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return open, nil
}
