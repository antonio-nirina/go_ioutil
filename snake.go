package main

import (
	//"bufio"
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"os/exec"
	"path/filepath"
	"strings"
	"time"
	//"runtime"

	// "github.com/dghubble/sling"
	"github.com/gofrs/uuid"
	// "github.com/gorilla/mux"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
	"golang.org/x/net/context"
	// "github.com/nsf/termbox-go"
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

type Env struct {
	FichierDonnees Info   `json:"fichier_donnees"`
	TypeEnveloppe  string `json:"type_enveloppe"`
}

type Info struct {
	Format        string `json:"format"`
	ContenuBase64 string `json:"contenu_base64"`
}

type Resp struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    []map[string]string `json:"data"`
}

func main25() {
	// printer()
	Init()
}

func HomeHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Println(ctx)
	env := Env{}
	resp := Resp{}
	err := json.NewDecoder(r.Body).Decode(&env)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	file, err := base64.StdEncoding.DecodeString(env.FichierDonnees.ContenuBase64)

	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	path := fmt.Sprintf("%s%s", filepath.Dir(""), "/csv/")
	filInfo, _ := os.Stat(path)

	if filInfo == nil {
		if err = os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			fmt.Println("contact_in_adminstartor1")
		}

		out, err := os.Create(path)

		if err != nil {
			fmt.Println(err)
		}
		defer out.Close()
		err = out.Chmod(0644)
	}

	u2, err := uuid.NewV4()

	if err != nil {
		fmt.Println("contact_in_adminstartor3")
	}
	name := fmt.Sprintf("%s%s%s", u2, ".", env.FichierDonnees.Format)
	filename := fmt.Sprintf("%s%s%s", path, "/", name)
	err = ioutil.WriteFile(filename, file, 0644)

	if err != nil {
		fmt.Println("contact_in_adminstartor4")
	}
	// Read CSV File
	csvData, err := ioutil.ReadFile("csv/" + name)

	if err != nil {
		panic(err)
	}

	bytes.Trim(csvData, "\xef\xbb\xbf")
	/*buff := bytes.NewReader(csvData)
	s := bufio.NewScanner(buff)
	var read_line string

	for s.Scan() {
		read_line = s.Text()
		read_line = strings.TrimSuffix(read_line, "\n")
	}
	copy(csvData, read_line)*/
	buff2 := bytes.NewReader(csvData)
	rec := csv.NewReader(buff2)
	rows := []map[string]string{}
	var header []string
	var headerLst []string
	var array []interface{}
	var res []interface{}
	var aHeader []interface{}

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
			// var dict []interface{}
			array = append(array, header)
		}
	}
	/*
		dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
	*/

	// fmt.Println("rrr", array[0])
	// fmt.Println("hhhh", array)
	// var dict map[string]interface{}

	// rows = append(rows, dict)
	for k, val := range array {
		if k == 0 {
			aHeader = append(aHeader, val)
		}
	}

	var in string
	var aIn string
	var az []interface{}

	for _, v := range aHeader {
		for _, val := range v.([]string) {
			in = val
		}
	}
	headerLst = strings.Split(in, ";")

	for key, rs := range array {
		if key > 0 {
			res = append(res, rs)
		}
	}

	for _, v1 := range res {
		for _, val1 := range v1.([]string) {
			aIn = val1
			az = append(az, strings.Split(aIn, ";"))
		}
	}

	/*dict := make(map[string]string)
	for _,v := range headerLst {
		for _,val2 := range az {
			dict[v] = val2
		}
	}
	rows = append(rows, dict)*/
	fmt.Println("hhhh", headerLst)
	// sc := az[0]

	for _, v21 := range az {
		for _, v22 := range v21.([]string) {
			fmt.Println("vvvv1", v22)
		}
	}

	if err != nil {
		log.Fatalf("r.Read() failed with '%s'\n", err)
	}

	resp.Code = 200
	resp.Message = "file is success store"
	resp.Data = rows
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

func Init() {
	// r := mux.NewRouter()
	r := xmux.New()
	c := xhandler.Chain{}
	r.POST("/", xhandler.HandlerFuncC(HomeHandler))
	fmt.Println("Run in 8080")
	// log.Fatal(http.ListenAndServe(":8080", r))
	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}

// We need cmd.Run for execute command and Stdout for output

// *********** Running a command and showing output **********************
// cmd.Stdoutet cmd.Stderrsont déclarés comme io.Writerinterface afin que
// nous puissions les définir sur n'importe quel type qui implémente une Write()méthode

// time.Date(time.Now().Year(), time.Now().Month())

func showTimes() {
	now := time.Now().Local()
	fmt.Println(now)
}

func converterClient() {
	fileName := "test_converter.docx"
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

/*func checkSl() {
	githubBase := sling.New().Base("https://api.github.com/").Client(httpClient)
	path := fmt.Sprintf("repos/%s/%s/issues", owner, repo)
	issues := new([]Issue)
	resp, err := githubBase.New().Get(path).QueryStruct(params).ReceiveSuccess(issues)
	fmt.Println(issues, resp, err)
}*/
