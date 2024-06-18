package configs

import "os"


type MinioConfigStruct struct {
	AccessKey string
	SecretKey string
	HostUrl string
	BlogBucket string
}

var MinioConfig MinioConfigStruct


func SetMinioConfig() {
	MinioConfig.AccessKey = os.Getenv("MINIO_ACCESSKEY")
	MinioConfig.HostUrl = os.Getenv("MINIO_URL")
	MinioConfig.SecretKey = os.Getenv("MINIO_SECRET")
	MinioConfig.BlogBucket = os.Getenv("MINIO_BLOG_BUCKET")
}
