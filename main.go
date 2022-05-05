package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"anime-skip.com/remote-config/src/backend"
	"anime-skip.com/remote-config/src/backend/aws"
	"anime-skip.com/remote-config/src/backend/cache"
	"anime-skip.com/remote-config/src/backend/env"
	"anime-skip.com/remote-config/src/backend/operations"
	"anime-skip.com/remote-config/src/backend/utils"
	"github.com/go-chi/chi"
)

func main() {
	s3Repo, err := aws.NewS3Repo(env.S3_BUCKET, env.S3_FILE_PATH)
	checkError(err)

	repo, err := cache.NewCacheRepo(s3Repo)
	checkError(err)

	router := backend.CreateRouter()

	ui := getUIFileSystem()
	err = operations.SetupUIRoutes(router, ui)
	checkError(err)

	// Unauthorized operations
	getHealth := operations.HealthHandler()
	getConfig := operations.GetAppConfigHandler(repo)
	// Authorized operations
	listApps := authMiddleware(operations.GetAppsHandler(repo))
	updateConfig := authMiddleware(operations.UpdateAppConfigHandler(repo))
	deleteConfig := authMiddleware(operations.DeleteAppConfigHandler(repo))

	router.Route("/api/", func(apiRoute chi.Router) {
		apiRoute.Get("/health", getHealth)
		apiRoute.Route("/config/{app}", func(configRoute chi.Router) {
			configRoute.Get("/", getConfig)
			configRoute.Put("/", updateConfig)
			configRoute.Delete("/", deleteConfig)
		})
		apiRoute.Get("/apps", listApps)
	})

	port := ":" + env.PORT
	fmt.Printf("Backend started  @ %s\n", port)
	http.ListenAndServe(port, router)
}

//go:embed dist/*
var prodUI embed.FS

func getUIFileSystem() *fs.FS {
	ui, err := fs.Sub(prodUI, "dist")
	if err != nil {
		panic("ui folder not found in embedded files")
	}
	return &ui
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func authMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := utils.CheckAuthHeader(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler(w, r)
	})
}
