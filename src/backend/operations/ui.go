package operations

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

func SetupUIRoutes(r *chi.Mux, website *fs.FS) {
	fileServer(r, "/", website)
}

func fileServer(r *chi.Mux, public string, assets *fs.FS) {
	index, err := fs.ReadFile(*assets, "index.html")
	if err != nil {
		panic(err)
	}
	embeddedFs := http.StripPrefix(public, http.FileServer(http.FS(*assets)))

	if strings.ContainsAny(public, "{}*") {
		panic("FileServer does not permit URL parameters.")
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
}
