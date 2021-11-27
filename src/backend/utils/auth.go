package utils

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"anime-skip.com/remote-config/src/backend/env"
)

func CheckAuthHeader(r *http.Request) error {
	auth := r.Header.Get("Authorization")
	expectedAuth := fmt.Sprintf("Bearer %s", env.AUTH_TOKEN)
	if auth != expectedAuth {
		time.Sleep(time.Second * 2)
		return errors.New("Unauthorized")
	}
	return nil
}
