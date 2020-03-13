package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main(){
	var folder []string
	folder = append(folder,"csv")
	folder = append(folder,"xcsv")
	
	for _,dirs := range folder {
		if err := os.MkdirAll(dirs, 0755); err != nil {
			fmt.Println(err)
		}
		content, err := ioutil.ReadDir("csv")
		
		if err != nil {
			fmt.Println(err)
		}

		for _, f := range content {
			info,err := os.Stat(f.Name())
			if err != nil {
				fmt.Println(err)
			}
        	fmt.Println(info)
    	}
	}
	
}