// package main

// import (
// 	"github.com/loheagn/gopic/cmd"
// )

// func main() {
// 	cmd.Execute()
// }

package main

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v6"
)

func main() {
	endpoint := "http://10.251.0.37:9000"
	accessKeyID := "admin"
	secretAccessKey := "2OqdKbvTUD8q2bOa"
	useSSL := false

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient初使化成功

	buckets, err := minioClient.ListBuckets()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}
