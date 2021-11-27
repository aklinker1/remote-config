package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"anime-skip.com/remote-config/src/backend"
	"anime-skip.com/remote-config/src/backend/aws"
	"anime-skip.com/remote-config/src/backend/cache"
	"anime-skip.com/remote-config/src/backend/operations"
	"github.com/go-chi/chi"
)

func main() {
	ui := getUIFileSystem()
	bucket := os.Getenv("S3_BUCKET")
	filePath := os.Getenv("S3_FILE_PATH")

	s3Repo, err := aws.NewS3Repo(bucket, filePath)
	checkError(err)

	repo, err := cache.NewCacheRepo(s3Repo)
	checkError(err)

	router := backend.CreateRouter()

	err = operations.SetupUIRoutes(router, ui)
	checkError(err)

	router.Route("/api/", func(apiRoute chi.Router) {
		apiRoute.Route("/config/{app}", func(configRoute chi.Router) {
			configRoute.Get("/", operations.GetAppConfigHandler(repo))
			configRoute.Put("/", operations.UpdateAppConfigHandler(repo))
			configRoute.Delete("/", operations.DeleteAppConfigHandler(repo))
		})
		apiRoute.Get("/apps", operations.GetAppsHandler(repo))
	})

	port := ":" + os.Getenv("PORT")
	fmt.Printf("Backend started  @ %s\n", port)
	http.ListenAndServe(port, router)
}

//go:embed src/backend/mock-ui/*
var devUI embed.FS

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
