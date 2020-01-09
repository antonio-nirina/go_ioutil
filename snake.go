package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	// "github.com/dghubble/sling"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
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

func main() {
	// printer()
	Init()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	env := Env{}
	// info := Info{}
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
	path := fmt.Sprintf("%s%s", filepath.Dir(""), "/fichier/")
	filInfo, _ := os.Stat(path)

	if filInfo == nil {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			fmt.Println("contact_in_adminstartor")
		}

		out, err := os.Create(path)

		if err != nil {
			fmt.Println("contact_in_adminstartor")
		}
		defer out.Close()
		err = out.Chmod(0644)
	}

	u2, err := uuid.NewV4()

	if err != nil {
		fmt.Println("contact_in_adminstartor")
	}

	filename := fmt.Sprintf("%s%s", u2, env.FichierDonnees.Format)
	f, err := os.OpenFile("fichier/"+filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		fmt.Println("contact_in_adminstartor")
	}
	// don't forget to close it
	defer f.Close()

	fmt.Println(file)
	//if os.IsNotExist(err) {

	// }

	/*response, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))*/
}

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("POST")
	fmt.Println("Run in 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func tesst() {
	ts := [...]float32{7, -10, 13, 8, 4, -7.2, -12, -3.7, 3.5, -9.6, 6.5, -1.7, -6.2, 7, 0.5, -0.3}
	var pos []float32
	var neg []float32

	for _, val := range ts {
		if val > 0 {
			pos = append(pos, val)
		} else {
			neg = append(neg, val)
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

func showTimes() {
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
	fmt.Println(string(out))
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
