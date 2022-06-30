package getpath

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var p PathModel

func Path(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &p)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(p)

	newData, err := json.Marshal(p)
	fmt.Println(string(newData))

	err = ioutil.WriteFile("path.json", newData, 0644)

	CreateFolder(p.Path)
}

func CreateFolder(path string) {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil && os.IsExist(err) {
		fmt.Println(err)
	} else {
		fmt.Println("Folder created")
	}
}
