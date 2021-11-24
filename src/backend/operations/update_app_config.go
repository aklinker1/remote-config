package operations

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"anime-skip.com/remote-config/src/backend"
	"github.com/go-chi/chi"
)

func UpdateAppConfigHandler(repo backend.Repo) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var newConfig backend.JSON
		err = json.Unmarshal(body, &newConfig)
		if err != nil {
			panic(err)
		}
		app := chi.URLParam(r, "app")
		println(app)

		repo.SaveConfig(app, newConfig)
		rw.WriteHeader(http.StatusNoContent)
	}
}
