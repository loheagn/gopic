package cmd

type S3Config struct {
	Endpoint  string `json:"endpoint"`
	Bucket    string `json:"bucket"`
	Accesskey string `json:"accesskey"`
	Secretkey string `json:"secretkey"`
	SSL       bool   `json:"ssl"`
}
