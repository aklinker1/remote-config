package operations

import (
	"net/http"

	"anime-skip.com/remote-config/src/backend"
	"anime-skip.com/remote-config/src/backend/utils"
)

func GetAppsHandler(repo backend.Repo) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		utils.SendJSON(rw, http.StatusOK, repo.GetApps())
	}
}
