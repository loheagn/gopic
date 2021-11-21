package cmd

import (
	"context"
	"strings"

	fp "path/filepath"

	"github.com/minio/minio-go/v6"
)

func s3Upload(ctx context.Context, conf *S3Config, filepath string) (string, error) {
	objName := func() string {
		slice := strings.Split(filepath, "/")
		return slice[len(slice)-1]
	}()
	filepath, err := fp.Abs(filepath)
	if err != nil {
		return "", err
	}

	minioClient, err := minio.New(conf.Endpoint, conf.Accesskey, conf.Secretkey, conf.SSL)
	if err != nil {
		return "", err
	}
	// TODO: check size
	_, err = minioClient.FPutObjectWithContext(ctx, conf.Bucket, objName, filepath, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return s3Filename(conf, objName), nil
}
