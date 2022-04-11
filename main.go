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

	router.Route("/api/", func(apiRoute chi.Router) {
		apiRoute.Get("/health", operations.HealthHandler())
		apiRoute.Route("/config/{app}", func(configRoute chi.Router) {
			configRoute.Get("/", operations.GetAppConfigHandler(repo))
			configRoute.Put("/", operations.UpdateAppConfigHandler(repo))
			configRoute.Delete("/", operations.DeleteAppConfigHandler(repo))
		})
		apiRoute.Get("/apps", operations.GetAppsHandler(repo))
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
