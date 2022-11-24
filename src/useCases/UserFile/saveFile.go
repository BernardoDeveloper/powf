package userfile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	compressfile "github.com/BernardoDeveloper/powf/src/useCases/CompressFile"
	getpath "github.com/BernardoDeveloper/powf/src/useCases/GetPath"
	"github.com/BernardoDeveloper/powf/src/utils"
)

func SaveFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	utils.HttpErr("Error retrievin the file.", err)

	// TODO: create automatized path
	content, _ := ioutil.ReadFile("/home/bernardo/programacao/powf/path.json")
	var p getpath.PathModel
	json.Unmarshal(content, &p)
	path := p.Path

	f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	utils.HttpErr("Can't open the file.", err)

	defer f.Close()
	io.Copy(f, file)

	// call compress function
	compressfile.CompressFileMethod(path, handler.Filename)

	fmt.Printf("Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("File size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
}
