package userfile

import (
	"fmt"
	"mime/multipart"
	"net/http"
)

func GetFile() (multipart.File, *multipart.FileHeader) {
	r := http.Request{}
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving the file.")
		panic(err)
	}

	return file, handler
}
