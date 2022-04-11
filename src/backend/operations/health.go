package operations

import (
	"net/http"

	"anime-skip.com/remote-config/src/backend"
	"anime-skip.com/remote-config/src/backend/env"
	"anime-skip.com/remote-config/src/backend/utils"
)

func HealthHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		utils.SendJSON(rw, http.StatusOK, backend.Health{
			Version: env.VERSION,
			Stage:   env.STAGE,
		})
	}
}
