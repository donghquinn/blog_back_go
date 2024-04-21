package database

import (
	"context"
	"log"

	"github.com/donghquinn/blog_back_go/configs"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)


func MinioInstance() (*minio.Client, error) {
	 minioConfig := configs.MinioConfig

	client, err := minio.New(minioConfig.HostUrl, &minio.Options{
		Creds: credentials.NewStaticV4(minioConfig.AccessKey, minioConfig.SecretKey, ""),
		Secure: true,
	})

	if err != nil {
		log.Printf("[MINIO] Initiate Minio Connection Instance Err: %v", err)

		return nil, err
	}

	return client, nil
}

func MinioConnect() error {
	minioConfig := configs.MinioConfig
	client, err := MinioInstance()

	if err != nil {
		return err
	}

	makeErr := client.MakeBucket(context.Background(), minioConfig.BlogBucket, minio.MakeBucketOptions{})

	if makeErr != nil {
		log.Printf("[MINIO] Making New Bucket Error: %v", makeErr)
		isExist, existErr := client.BucketExists(context.Background(), minioConfig.BlogBucket)

		if isExist && existErr == nil {
			log.Printf("[MINIO] Found Already Existing Bucket")
		} else {
			log.Printf("[MINIO] Check Minio Bucket Error: %v", existErr)

			return existErr
		}
	} else {
		log.Printf("[MINIO] Successfully Created")
	}


	return nil
}

func UploadVideo(objectName string, filePath string, videoType string) (minio.UploadInfo,error) {
	minioConfig := configs.MinioConfig
	minioClient, minioErr := MinioInstance()

	if minioErr != nil {
		log.Printf("[UPLOAD] Connect Minio Error: %v", minioErr)

		return minio.UploadInfo{}, minioErr
	}

	uploadInfo, uploadErr := minioClient.FPutObject(context.Background(), minioConfig.BlogBucket, objectName, filePath, minio.PutObjectOptions{ContentType: videoType})

	if uploadErr != nil {
		log.Printf("[UPLOAD] Upload Video Error: %v", uploadErr)

		return minio.UploadInfo{}, uploadErr
	}
	
	return uploadInfo, nil
}

func UploadImage(objectName string, filePath string, videoType string) (minio.UploadInfo,error) {
	minioConfig := configs.MinioConfig
	minioClient, minioErr := MinioInstance()

	if minioErr != nil {
		log.Printf("[UPLOAD] Connect Minio Error: %v", minioErr)

		return minio.UploadInfo{}, minioErr
	}

	uploadInfo, uploadErr := minioClient.FPutObject(context.Background(), minioConfig.BlogBucket, objectName, filePath, minio.PutObjectOptions{ContentType: videoType})

	if uploadErr != nil {
		log.Printf("[UPLOAD] Upload Image Error: %v", uploadErr)

		return minio.UploadInfo{}, uploadErr
	}
	
	return uploadInfo, nil
}

func GetVideo(objectName string) (*minio.Object, error) {
	minioConfig := configs.MinioConfig
	minioClient, minioErr := MinioInstance()

		if minioErr != nil {
		log.Printf("[STREAM] Connect Minio Error: %v", minioErr)

		return nil, minioErr
	}

	video, getErr := minioClient.GetObject(context.Background(), minioConfig.BlogBucket, objectName, minio.GetObjectOptions{} )

	if getErr != nil {
		log.Printf("[STREAM] Get Video Error: %v", getErr)

		return nil, getErr
	}

	// https://stackoverflow.com/questions/73669999/golang-minio-client-can-put-and-remove-but-not-stat-or-get-objects
	return video, nil
}