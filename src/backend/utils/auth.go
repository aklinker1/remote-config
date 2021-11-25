package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

func CheckAuthHeader(r *http.Request) error {
	auth := r.Header.Get("Authorization")
	expectedAuth := fmt.Sprintf("Bearer %s", os.Getenv("AUTH_TOKEN"))
	if auth != expectedAuth {
		time.Sleep(time.Second * 2)
		return errors.New("Unauthorized")
	}
	return nil
}
