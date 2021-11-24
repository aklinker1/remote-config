package utils

import (
	"net/http"
	"strings"
)

func AppFromConfigRequest(r *http.Request) string {
	suffix := strings.TrimPrefix(r.URL.Path, "/api/config/")
	return strings.TrimSuffix(suffix, "/")
}
