package backend

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors)
	r.Use(middleware.Recoverer)
	return r
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", strings.Join([]string{
			"Accept",
			"Authorization",
			"Content-Type",
		}, ","))
		res.Header().Set("Access-Control-Allow-Methods", strings.Join([]string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		}, ","))

		if req.Method == http.MethodOptions {
			res.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(res, req)
		}
	})
}
