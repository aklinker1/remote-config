package cache

import (
	"anime-skip.com/remote-config/src/backend"
)

type cacheRepo struct {
	blob    backend.Blob
	wrapped backend.Repo
}

func NewCacheRepo(wrapped backend.Repo) (*cacheRepo, error) {
	initialBlob, err := wrapped.GetBlob()
	return &cacheRepo{
		blob:    initialBlob,
		wrapped: wrapped,
	}, err
}

func (cache *cacheRepo) GetBlob() (backend.Blob, error) {
	return cache.blob, nil
}

func (cache *cacheRepo) GetApps() ([]string, error) {
	apps := []string{}
	blob, err := cache.GetBlob()
	if err != nil {
		return nil, err
	}
	for app := range blob {
		apps = append(apps, app)
	}
	return apps, nil
}

func (cache *cacheRepo) GetConfig(app string) (backend.JSON, error) {
	blob, err := cache.GetBlob()
	if err != nil {
		return nil, err
	}
	config, exists := blob[app]
	if exists {
		return config, nil
	}
	return backend.JSON{}, nil
}

func (cache *cacheRepo) SaveConfig(app string, config backend.JSON) (backend.JSON, error) {
	savedConfig, err := cache.wrapped.SaveConfig(app, config)
	if err == nil {
		// TODO: Mutex read/write of cache?
		cache.blob[app] = savedConfig
	}
	return savedConfig, err
}

func (cache *cacheRepo) DeleteConfig(app string) (backend.JSON, error) {
	config, err := cache.wrapped.DeleteConfig(app)
	if err == nil {
		// TODO: Mutex read/write of cache?
		delete(cache.blob, app)
	}
	return config, err
}
