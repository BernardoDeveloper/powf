package utils

import (
	"fmt"
	"os"
)

func HttpErr(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stdout, "\033[0;31m [ERROR]: { \033[0m %s \033[0;31m } \033[0m \n \033[0;31m [PANIC MSG]: \033[0m", msg)
		panic(err)
	}
}
