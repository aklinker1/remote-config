package utils

import (
	"encoding/json"
	"net/http"

	"anime-skip.com/remote-config/src/backend"
)

func SendJSON(res http.ResponseWriter, status int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	res.Write(bytes)
}

func SendErrorJSON(res http.ResponseWriter, err error) {
	SendJSON(res, http.StatusInternalServerError, backend.JSON{
		"error": err.Error(),
	})
}
