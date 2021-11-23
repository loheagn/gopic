package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	fp "path/filepath"

	"github.com/minio/minio-go/v6"
)

func s3Upload(ctx context.Context, conf *S3Config, filepath string) (string, error) {
	objName := func() string {
		slice := strings.Split(filepath, "/")
		return slice[len(slice)-1]
	}()

	conf.Endpoint = func(endpoint string) string {
		endpoint = strings.TrimPrefix(endpoint, "http://")
		endpoint = strings.TrimPrefix(endpoint, "https://")
		return endpoint
	}(conf.Endpoint)

	filepath, err := fp.Abs(filepath)
	if err != nil {
		return "", err
	}
	// check filepath
	localfile, err := os.Stat(filepath)
	if err != nil {
		return "", err
	}
	if localfile.IsDir() {
		return "", fmt.Errorf("can not upload a dir: %s", filepath)
	}

	minioClient, err := minio.New(conf.Endpoint, conf.Accesskey, conf.Secretkey, conf.SSL)
	if err != nil {
		return "", err
	}

	n, err := minioClient.FPutObjectWithContext(ctx, conf.Bucket, objName, filepath, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	if n != localfile.Size() {
		return "", fmt.Errorf("upload failed: %s", filepath)
	}

	return s3Filename(conf, objName), nil
}
