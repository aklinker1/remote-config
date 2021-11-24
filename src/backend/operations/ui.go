package operations

import (
	"errors"
	"io/fs"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

func SetupUIRoutes(r *chi.Mux, website *fs.FS) error {
	return fileServer(r, "/", website)
}

func fileServer(r *chi.Mux, public string, assets *fs.FS) error {
	index, err := fs.ReadFile(*assets, "index.html")
	if err != nil {
		return err
	}
	embeddedFs := http.StripPrefix(public, http.FileServer(http.FS(*assets)))

	if strings.ContainsAny(public, "{}*") {
		return errors.New("FileServer does not permit URL parameters.")
	}

	if public != "/" && public[len(public)-1] != '/' {
		r.Get(public, http.RedirectHandler(public+"/", 301).ServeHTTP)
		public += "/"
	}

	r.Get(public+"*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.ContainsRune(r.URL.Path, '.') {
			w.WriteHeader(200)
			w.Header().Add("Content-Type", "text/html")
			w.Write(index)
		} else {
			embeddedFs.ServeHTTP(w, r)
		}
	}))

	return nil
}
