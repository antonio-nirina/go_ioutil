package main

import (
	"fmt"
	"os/exec"
	//"os"
)

// go build -o E:\project\Go\bin ## move executable in bin folder

func cmd() {
	// InitFile()
	// fmt.Println("file_created")
	cmd := exec.Command("unoconv")
	fmt.Println(cmd)
}
