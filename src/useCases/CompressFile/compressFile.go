package compressfile

import (
	"bufio"
	"compress/gzip"
	"io/ioutil"
	"os"
	"strings"

	userfile "github.com/BernardoDeveloper/powf/src/useCases/UserFile"
)

func CompressFileMethod() {
	var extension_file string
	_, handler := userfile.GetFile()

	f, _ := os.Open(handler.Filename)
	read := bufio.NewReader(f)
	data, _ := ioutil.ReadAll(read)

	handler.Filename = strings.Replace(handler.Filename, extension_file, ".gz", -1)
	f, _ = os.Create(handler.Filename)

	w := gzip.NewWriter(f)
	w.Write(data)
	w.Close()
}
