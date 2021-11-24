package operations

import (
	"net/http"

	"anime-skip.com/remote-config/src/backend"
)

func NotFoundHandler(repo backend.Repo) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	}
}
