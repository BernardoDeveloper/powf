package userfile

import (
	"fmt"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// file, _ := GetFile()
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving the file.")
		panic(err)
	}

	/* Create temp file
	tempFile, err := ioutil.TempFile("file", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()*/

	// Read all content upload file into byte array
	/*fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}*/

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// tempFile.Write(fileBytes)
	fmt.Fprintf(w, "successfully uploaded file\n")
}
