package aws

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"anime-skip.com/remote-config/src/backend"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Repo struct {
	s3       *s3.Client
	bucket   *string
	filename *string
}

func NewS3Repo(bucket string, filename string) (*s3Repo, error) {
	ctx := context.TODO()

	// Get the client
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)

	repo := &s3Repo{
		s3:       client,
		bucket:   &bucket,
		filename: &filename,
	}

	// Make sure the bucket exists
	initialData, err := repo.GetBlob()
	if err != nil {
		return nil, err
	}
	fmt.Println("Initialized blob to", initialData)

	return repo, nil
}

func (repo *s3Repo) GetBlob() (backend.Blob, error) {
	object, err := repo.s3.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: repo.bucket,
		Key:    repo.filename,
	})
	if err != nil {
		return nil, err
	}
	if *object.ContentType != "application/json" {
		return nil, fmt.Errorf("Only JSON data is supported, was %s", *object.ContentType)
	}
	data, err := ioutil.ReadAll(object.Body)
	if err != nil {
		return nil, err
	}
	var blob backend.Blob
	err = json.Unmarshal(data, &blob)
	if err != nil {
		return nil, err
	}
	return blob, nil
}

func (repo *s3Repo) GetApps() ([]string, error) {
	panic("Not implemented: s3Repo.GetApps")
}

func (repo *s3Repo) GetConfig(app string) (backend.JSON, error) {
	panic("Not implemented: s3Repo.GetApps")
}

func (repo *s3Repo) writeBlob(blob backend.Blob) error {
	body, err := json.Marshal(blob)
	if err != nil {
		return err
	}
	repo.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      repo.bucket,
		Key:         repo.filename,
		Body:        bytes.NewReader(body),
		ContentType: aws.String("application/json"),
	})
	return nil
}

func (repo *s3Repo) SaveConfig(app string, config backend.JSON) (backend.JSON, error) {
	blob, err := repo.GetBlob()
	if err != nil {
		return nil, err
	}

	blob[app] = config
	return config, repo.writeBlob(blob)
}

func (repo *s3Repo) DeleteConfig(app string) (backend.JSON, error) {
	blob, err := repo.GetBlob()
	if err != nil {
		return nil, err
	}

	config := blob[app]
	delete(blob, app)
	return config, repo.writeBlob(blob)
}
