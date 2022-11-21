package userfile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	getpath "github.com/BernardoDeveloper/powf/src/useCases/GetPath"
)

func SaveFile(w http.ResponseWriter, r *http.Request) {
	//file, handler := GetFile()
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving the file.")
		panic(err)
	}

	// TODO: create automatized path
	content, _ := ioutil.ReadFile("/home/bernardo/programacao/powf/path.json")
	var p getpath.PathModel
	json.Unmarshal(content, &p)
	path := p.Path

	f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, file)

	// defer file.Close()
	fmt.Printf("Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("File size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
}
