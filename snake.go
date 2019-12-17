package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"log"
	"os/exec"
	"io/ioutil"

	"github.com/gofrs/uuid"
)

const (
	URL_S3 = "https://s3.eu-west-2.amazonaws.com/"
	URL_SERVER = "http://test-apiconverter.servicepostal.com/file/convert"
)

type  Inputs struct {
	File Model `json:"file"`
	UploadType int `json:"upload_type"`
}

type Model struct {
	FileName string `json:"file_name"`
	FileUrl string `json:"file_url"`
}

func main(){
	// printer()
	converterClient()
}

func printer(){
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 50; j++ {
			if j > i {
				for k := 1; k <= i; k ++ {
					fmt.Print("*")
				}
				break		
			} else {
				fmt.Print("*")
			}
		}
		
		fmt.Println("")	
	}

	for m := 1; m <= 5; m++ {
		for n := 1; n <= 25; n++ {
			if n >= 20 {
				fmt.Print("*")	
			} else {
				fmt.Print(" ")		
			}
		}
		
		fmt.Println("")	
	}
}

// We need cmd.Run for execute command and Stdout for output

// *********** Running a command and showing output **********************
// cmd.Stdoutet cmd.Stderrsont déclarés comme io.Writerinterface afin que 
// nous puissions les définir sur n'importe quel type qui implémente une Write()méthode

// time.Date(time.Now().Year(), time.Now().Month())

func monitor() {
	cmd := exec.Command("ls","-a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func converterClient() {
	fileName := "067b77b3-c18e-4bd0-b4f1-e7e63020812c16-12-2019.docx"
	bucket := "sp-files-to-convert"
	method := "POST"
	uri := fmt.Sprintf("%s%s%s%s",URL_S3,bucket,"/",fileName)
	u2, err := uuid.NewV4()

	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	name := fmt.Sprintf("%s%s",u2,".docx")
	payload := Inputs{
		File:Model{
			FileName:name,
			FileUrl:uri,
		},
		UploadType: 2,
	}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(payload)
	req, err := http.NewRequest(method, URL_SERVER, reqBodyBytes)
	
	if err != nil {
		log.Fatalf("creating request: %s\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp,err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error_do_req %v", err)
		log.Println("Body:", err)
		// return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
