package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main(){
	var folder []string
	var res [][]os.FileInfo
	var list = make(map[string]interface{})
	var result []interface{}
	folder = append(folder,"csv")
	folder = append(folder,"xcsv")
	
	for _,dirs := range folder {
		if err := os.MkdirAll(dirs, 0755); err != nil {
			fmt.Println("errr000",err)
		}
		content, err := ioutil.ReadDir(dirs)
		
		if err != nil {
			fmt.Println("errr22",err)
		}
		res = append(res,content)
	}

	for _, f := range res {
		for _, val := range f {
			list["name"] = val.Name()
			list["date"] = val.ModTime()
			result = append(result,list)
	    }
    }
    fmt.Println("info", result)
	
}