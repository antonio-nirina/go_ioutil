package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func InitFile() {
	// path := "/home/nirina/Documents/golang/example/var/"
	path := fmt.Sprintf("%s%s", filepath.Dir(""), "/var/") 
	if !fileExists(path) {
		mkdirForFile(path)	
	}
	createFile()
	// fmt.Println(path)
}

// Base returns the file name after the last slash.
// Dir returns the directory without the last file name
func mkdirForFile(fpath string) error {
	if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
		return fmt.Errorf("%s: making directory for file: %v", fpath, err)
	}
	out, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("%s: creating new file: %v", fpath, err)
	}
	defer out.Close() // nolint: errcheck
	err = out.Chmod(0644)

	if err != nil && runtime.GOOS != "windows" {
		return fmt.Errorf("%s: changing file mode: %v", fpath, err)
	}
	/*
		copy File
		_, err = io.Copy(out, in)
		if err != nil {
			return fmt.Errorf("%s: writing file: %v", fpath, err)
		}
	*/
	return nil
}

//===================================
// fileExists
//===================================
func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// CreateFile in directory
func createFile () {
	f, err := os.OpenFile("var/info.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	// don't forget to close it
	defer f.Close()
}