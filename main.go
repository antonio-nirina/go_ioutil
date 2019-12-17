package main

import (
	"fmt"
	"os/exec"
	//"os"
)

func cmd() {
	// InitFile()
	// fmt.Println("file_created")
	cmd := exec.Command("unoconv")
	fmt.Println(cmd)
}

