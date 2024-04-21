package configs

import (
	"fmt"
	"os"
)

type GlobalConfigStruct struct {
	AppPort string
	AppHost string
	AesIv string
	AesKey string
	SecretKey string
	JwtKey string
}

type DatabaseConfigStruct struct {
	Host string
	Port string
	Database string
	User string
	Password string
}

var GlobalConfig GlobalConfigStruct
var MinioConfig MinioConfigStruct
var DatabaseConfig DatabaseConfigStruct

type MinioConfigStruct struct {
	AccessKey string
	SecretKey string
	HostUrl string
	BlogBucket string
}

func SetGlobalConfig(){
	GlobalConfig.AppPort = os.Getenv("APP_PORT")
	GlobalConfig.AppHost = fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	GlobalConfig.AesIv = os.Getenv("AES_IV")
	GlobalConfig.AesKey = os.Getenv("AES_KEY")
	GlobalConfig.JwtKey = os.Getenv("JWT_KEY")
	GlobalConfig.SecretKey = os.Getenv("SECRET_KEY")
}

func SetDatabaseConfig() {
	DatabaseConfig.Host = os.Getenv("MARIADB_HOST")
	DatabaseConfig.User = os.Getenv("MARIADB_USER")
	DatabaseConfig.Password = os.Getenv("MARIADB_PASSWORD")
	DatabaseConfig.Port = os.Getenv("MARIADB_PORT")
	DatabaseConfig.Database = os.Getenv("MARIADB_DATABASE")
}

func SetMinioConfig() {
	MinioConfig.AccessKey = os.Getenv("MINIO_ACCESSKEY")
	MinioConfig.HostUrl = os.Getenv("MINIO_URL")
	MinioConfig.SecretKey = os.Getenv("MINIO_SECRET")
	MinioConfig.BlogBucket = os.Getenv("MINIO_BLOG_BUCKET")
}