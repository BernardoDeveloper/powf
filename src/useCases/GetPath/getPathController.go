package getpath

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type p struct {
// 	Path string `json:"path"`
// }

var p PathModel

func Serialize(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}

	fmt.Print(string(val))
	return string(val), nil
}

func Path(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil || r.ContentLength == 0 {
		panic("Request body is empty")
	}

	value := json.NewDecoder(r.Body).Decode(&p)
	result, _ := Serialize(value)

	// path := json.NewDecoder(r.Body).Decode(&p{})
	// if path == nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	fmt.Print(result)
	// os.Setenv("PATH", path)
}
