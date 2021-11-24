package utils

import (
	"encoding/json"
	"net/http"
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
