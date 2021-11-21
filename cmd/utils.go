package cmd

import "fmt"

func s3Filename(conf *S3Config, objName string) string {
	protocol := "http"
	if conf.SSL {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s/%s/%s", protocol, conf.Endpoint, conf.Bucket, objName)
}
