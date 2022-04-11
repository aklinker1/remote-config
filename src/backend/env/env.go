package env

import (
	"errors"
	"os"
)

// Environment Variables

var S3_BUCKET = os.Getenv("S3_BUCKET")
var S3_FILE_PATH = os.Getenv("S3_FILE_PATH")
var PORT = os.Getenv("PORT")
var AUTH_TOKEN = os.Getenv("AUTH_TOKEN")
var BUG_REPORTER = os.Getenv("BUG_REPORTER")

// Compile Time Variables

var (
	STAGE   string
	VERSION string
)

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
	if STAGE == "" {
		panic(errors.New("env.STAGE must be set at compile-time"))
	}
	if VERSION == "" {
		panic(errors.New("env.VERSION must be set at compile-time"))
	}
}
