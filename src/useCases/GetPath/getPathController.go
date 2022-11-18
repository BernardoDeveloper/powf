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

	newData, _ := json.Marshal(p)
	fmt.Println(string(newData))

	CreateFolder(p.Path)

	ioutil.WriteFile("path.json", newData, 0644)
}

func CreateFolder(path string) {
	if path == "" {
		fmt.Println("Can't path null")
		return
	}

	err := os.Mkdir(path, os.ModePerm)
	if err != nil && os.IsExist(err) {
		fmt.Println(err)
	} else {
		fmt.Println("Folder created")
	}
}
