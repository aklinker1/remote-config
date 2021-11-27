package env

import (
	"errors"
	"os"
)

var S3_BUCKET = os.Getenv("S3_BUCKET")
var S3_FILE_PATH = os.Getenv("S3_FILE_PATH")
var PORT = os.Getenv("PORT")
var AUTH_TOKEN = os.Getenv("AUTH_TOKEN")

func init() {
	if S3_BUCKET == "" {
		panic(errors.New("env.S3_BUCKET must be set"))
	}
	if S3_FILE_PATH == "" {
		panic(errors.New("env.S3_FILE_PATH must be set"))
	}
	if PORT == "" {
		panic(errors.New("env.PORT must be set"))
	}
	if AUTH_TOKEN == "" {
		panic(errors.New("env.AUTH_TOKEN must be set"))
	}
}
