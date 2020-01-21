package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"io/ioutil"
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

func main() {
	// csvfile, err := ioutil.ReadFile("csv/23680732-e3dd-40fb-a56c-562ce78a9de6.csv")
	csvfile, err := os.Open("csv/23680732-e3dd-40fb-a56c-562ce78a9de6.csv")
	if err != nil {
		fmt.Println("contact_in_adminstartor1")
	}

	r := strings.NewReader(csvfile) // XXX: open the file instead...
	rec := csv.NewReader(TrimReader{r})

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
