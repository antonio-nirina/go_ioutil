package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
	"regexp"
	// "strings"

	"github.com/rs/xhandler"
	"github.com/rs/xmux"
	"golang.org/x/net/context"
)

type TrimReader struct{ io.Reader }
type Resp1 struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    []map[string]string `json:"data"`
}

var trailingws = regexp.MustCompile(` +\r?\n`)

func (tr TrimReader) Read(bs []byte) (int, error) {
	n, err := tr.Reader.Read(bs)
	if err != nil {
		return n, err
	}
	line := string(bs[:n])
	trimmed := []byte(trailingws.ReplaceAllString(line, "\n"))
	copy(bs, trimmed)
	return len(trimmed), nil
}

func main() {
	Init1()
}

func HomeHandle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Println(ctx)
	csvfile, err := ioutil.ReadFile("csv/517e1951-628d-4cc1-8a3d-1eb4d187c62f.csv")
	if err != nil {
		fmt.Println("contact_in_adminstartor1")
	}

	resp := Resp1{}
	buff2 := bytes.NewReader(csvfile)
	csv := CSVToMap(buff2)
	fmt.Println("rrr", csv)

	resp.Code = 200
	resp.Message = "file is success store"
	resp.Data = csv
	response, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	w.Write([]byte(response))
}

func CSVToMap(reader io.Reader) []map[string]string {
	r := csv.NewReader(reader)
	rows := []map[string]string{}
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(header) == 0 {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}
	return rows
}

func Init1() {
	// r := mux.NewRouter()
	r := xmux.New()
	c := xhandler.Chain{}
	r.POST("/", xhandler.HandlerFuncC(HomeHandle))
	fmt.Println("Run in 8080")
	// log.Fatal(http.ListenAndServe(":8080", r))
	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}
