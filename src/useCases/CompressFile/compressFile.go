package compressfile

import (
	"bufio"
	"compress/gzip"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BernardoDeveloper/powf/src/utils"
)

func CompressFileMethod(path string, fileName string) {
	file, err := os.Open(path + fileName)
	utils.HttpErr("Can't open the file.", err)

	read := bufio.NewReader(file)
	data, err := ioutil.ReadAll(read)
	utils.HttpErr("Error to read file.", err)

	extension := filepath.Ext(path)
	compressFileName := strings.Replace(path, extension, ".gz", -1)

	file, err = os.Create(path + compressFileName)
	utils.HttpErr("Error to compress file.", err)

	writer := gzip.NewWriter(file)
	writer.Write(data)
	writer.Close()
}
