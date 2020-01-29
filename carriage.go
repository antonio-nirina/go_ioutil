package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	//"os"
	"regexp"
	"strings"
)

type TrimReader struct{ io.Reader }

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

func main12() {
	csvfile, err := ioutil.ReadFile("csv/a158cc8e-71bd-4edb-bdb6-d00e1b65ba25.csv")
	if err != nil {
		fmt.Println("contact_in_adminstartor1")
	}

	r := strings.NewReader(string(csvfile))
	rec := csv.NewReader(TrimReader{r})
	rows := []map[string]string{}
	var header []string

	for {
		record, err := rec.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		header = record

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
	fmt.Println("rrr", header)
	fmt.Println("hhhh", rows)
}
