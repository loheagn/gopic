package cmd

import (
	"context"
	"fmt"
	"testing"

	"github.com/jinzhu/copier"
)

// use minio-local-server for test
const (
	s3Endpoint  = "localhost:9000"
	s3Accesskey = "minioadmin"
	s3Secretkey = "minioadmin"
)

func Test_s3Upload(t *testing.T) {
	bucket := "test-bucket"
	s3Config := &S3Config{
		Endpoint:  s3Endpoint,
		Bucket:    bucket,
		Accesskey: s3Accesskey,
		Secretkey: s3Secretkey,
		SSL:       false,
	}
	s3ConfigWithProtocol := &S3Config{}
	err := copier.Copy(s3ConfigWithProtocol, s3Config)
	if err != nil {
		panic(err.Error())
	}
	s3ConfigWithProtocol.Endpoint = "http://" + s3Endpoint
	type args struct {
		ctx      context.Context
		conf     *S3Config
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "endpoint with protocol",
			args: args{
				ctx:      context.Background(),
				conf:     s3ConfigWithProtocol,
				filepath: "../test/upload_test.txt",
			},
			want: fmt.Sprintf("http://%s/%s/upload_test.txt", s3Endpoint, bucket),
		},
		{
			name: "endpoint without protocol",
			args: args{
				ctx:      context.Background(),
				conf:     s3Config,
				filepath: "../test/upload_test.txt",
			},
			want: fmt.Sprintf("http://%s/%s/upload_test.txt", s3Endpoint, bucket),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s3Upload(tt.args.ctx, tt.args.conf, tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("s3Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("s3Upload() got = %v, want %v", got, tt.want)
			}
		})
	}
}
