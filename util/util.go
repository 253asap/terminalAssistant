package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ErrorCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("Error: %s\n%v", msg, err)
		os.Exit(1)
	}
}

func GetEnvVar(envv string) (envVar string) {
	error := godotenv.Load()
	ErrorCheck("Couldnt load .env", error)
	return os.Getenv(envv)
}
