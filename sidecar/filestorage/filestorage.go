package filestorage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go"
	log "github.com/sirupsen/logrus"
	"io"
	"strings"
)

const ENV_PROD = "prod"

type UploadInput struct {
	File        io.Reader
	Name        string
	Size        int64
	ContentType string
}

type FileStorage struct {
	client   *minio.Client
	bucket   string
	endpoint string
	env      string
}

func NewFileStorage(client *minio.Client, bucket, endpoint, env string) *FileStorage {
	return &FileStorage{
		client:   client,
		bucket:   bucket,
		endpoint: endpoint,
		env:      env,
	}
}

func (fs *FileStorage) Upload(ctx context.Context, input UploadInput) (string, error) {
	opts := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	_, err := fs.client.PutObjectWithContext(ctx,
		fs.bucket, input.Name, input.File, input.Size, opts)
	if err != nil {
		log.Errorf("error occured while uploading file to bucket: %s", err.Error())
		return "", err
	}

	return fs.generateFileURL(input.Name), nil
}

func (fs *FileStorage) generateFileURL(fileName string) string {
	// DigitalOcean Spaces link format
	if fs.env == ENV_PROD {
		return fmt.Sprintf("https://%s.%s/%s", fs.bucket, fs.endpoint, fileName)
	}

	// localstack S3 link format
	endpoint := strings.Replace(fs.endpoint, "localstack", "localhost", -1)
	return fmt.Sprintf("http://%s/%s/%s", endpoint, fs.bucket, fileName)
}
