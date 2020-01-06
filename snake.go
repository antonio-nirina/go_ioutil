package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/dghubble/sling"
	"github.com/gofrs/uuid"
)

const (
	URL_S3     = "https://s3.eu-west-2.amazonaws.com/"
	URL_SERVER = "http://test-apiconverter.servicepostal.com/file/convert"
)

type Inputs struct {
	File       Model `json:"file"`
	UploadType int   `json:"upload_type"`
}

type Model struct {
	FileName string `json:"file_name"`
	FileUrl  string `json:"file_url"`
}

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	// printer()
	// converterClient()
	tesst()
}


func tesst() {
	ts := [...]float32{7,-10,13,8,4,-7.2,-12,-3.7,3.5,-9.6,6.5,-1.7,-6.2,7,0.5}
	var max float32 = ts[0]
	var min float32 = ts[0]
	var pos  []float32
	var neg  []float32

	for _, value := range ts {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }

	for _,val := range ts{
		
		if val <= max && val >= min {
	        if val > 0 {
	            pos = append(pos,val)
	        } else {
	            neg = append(neg,val)
	        }
    	}
	}

	minPos := pos[0]
	maxNeg := neg[0]

	for _, valp := range pos {
        if minPos > valp {
            minPos = valp
        }
    }

    for _, valn := range neg {
        if maxNeg < valn {
            maxNeg = valn
        }
    }

    if minPos > (-1)*maxNeg {
    	fmt.Println(maxNeg)
    } else {
    	fmt.Println(minPos)
    }
}

func printer() {
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 50; j++ {
			if j > i {
				for k := 1; k <= i; k++ {
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

func showTimes(){
	now := time.Now().Local()
	fmt.Println(now)
}

func monitor() {
	cmd := exec.Command("pwd")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(out1))
}

func converterClient() {
	fileName := "067b77b3-c18e-4bd0-b4f1-e7e63020812c16-12-2019.docx"
	bucket := "sp-files-to-convert"
	method := "POST"
	uri := fmt.Sprintf("%s%s%s%s", URL_S3, bucket, "/", fileName)
	u2, err := uuid.NewV4()

	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	name := fmt.Sprintf("%s%s", u2, ".docx")
	payload := Inputs{
		File: Model{
			FileName: name,
			FileUrl:  uri,
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
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error_do_req %v", err)
		log.Println("Body:", err)
		// return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func checkSl() {
	githubBase := sling.New().Base("https://api.github.com/").Client(httpClient)
	path := fmt.Sprintf("repos/%s/%s/issues", owner, repo)
	issues := new([]Issue)
	resp, err := githubBase.New().Get(path).QueryStruct(params).ReceiveSuccess(issues)
	fmt.Println(issues, resp, err)
}
