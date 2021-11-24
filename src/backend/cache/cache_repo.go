package cache

import (
	"anime-skip.com/remote-config/src/backend"
)

type cacheRepo struct {
	blob    backend.Blob
	wrapped backend.Repo
}

func NewCacheRepo(wrapped backend.Repo) *cacheRepo {
	return &cacheRepo{
		blob:    wrapped.GetBlob(),
		wrapped: wrapped,
	}
}

func (cache *cacheRepo) GetBlob() backend.Blob {
	return cache.blob
}

func (cache *cacheRepo) GetApps() []string {
	apps := []string{}
	for app := range cache.GetBlob() {
		apps = append(apps, app)
	}
	return apps
}

func (cache *cacheRepo) GetConfig(app string) backend.JSON {
	config, exists := cache.GetBlob()[app]
	if exists {
		return config
	}
	return backend.JSON{}
}

func (cache *cacheRepo) SaveConfig(app string, config backend.JSON) error {
	// TODO: Mutex read/write of cache?
	cache.blob[app] = config
	return cache.wrapped.SaveConfig(app, config)
}
