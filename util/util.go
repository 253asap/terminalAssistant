package util

import (
	"fmt"
	"os"
)

func ErrorCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("Error: %s\n%v", msg, err)
		os.Exit(1)
	}
}
