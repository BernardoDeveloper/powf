package utils

import (
	"fmt"
)

func HttpErr(msg string, err error) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}
