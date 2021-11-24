package aws

import (
	"anime-skip.com/remote-config/src/backend"
)

type s3Repo struct {
	blob    backend.Blob
	wrapped backend.Repo
}

func NewS3Repo() *s3Repo {
	return &s3Repo{}
}

func (s3 *s3Repo) GetBlob() backend.Blob {
	return backend.Blob{
		"test": {
			"some-value": true,
		},
	}
}

func (s3 *s3Repo) GetApps() []string {
	panic("Not implemented: s3Repo.GetApps")
}

func (s3 *s3Repo) GetConfig(app string) backend.JSON {
	panic("Not implemented: s3Repo.GetApps")
}

func (s3 *s3Repo) SaveConfig(app string, config backend.JSON) error {
	return nil
}
