package operations

import (
	"net/http"

	"anime-skip.com/remote-config/src/backend"
	"anime-skip.com/remote-config/src/backend/utils"
	"github.com/go-chi/chi"
)

func GetAppConfigHandler(repo backend.Repo) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		app := chi.URLParam(r, "app")

		config, err := repo.GetConfig(app)
		if err != nil {
			utils.SendErrorJSON(rw, err)
			return
		}
		utils.SendJSON(rw, http.StatusOK, config)
	}
}
