package operations

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"anime-skip.com/remote-config/src/backend"
	"anime-skip.com/remote-config/src/backend/utils"
	"github.com/go-chi/chi"
)

func UpdateAppConfigHandler(repo backend.Repo) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := utils.CheckAuthHeader(r)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var newConfig backend.JSON
		err = json.Unmarshal(body, &newConfig)
		if err != nil {
			utils.SendErrorJSON(rw, err)
			return
		}

		app := chi.URLParam(r, "app")
		savedConfig, err := repo.SaveConfig(app, newConfig)
		if err != nil {
			utils.SendErrorJSON(rw, err)
			return
		}

		utils.SendJSON(rw, http.StatusOK, savedConfig)
	}
}
