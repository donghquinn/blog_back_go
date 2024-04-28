package main

import (
	"log"
	"net/http"
	"time"

	"github.com/donghquinn/blog_back_go/configs"
	"github.com/donghquinn/blog_back_go/libraries/database"
	"github.com/donghquinn/blog_back_go/middlewares"
	"github.com/donghquinn/blog_back_go/routers"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Printf("[ENV] Load Env Error")
	}

	setConfigs()

	minioErr := database.MinioConnect()
	if minioErr != nil {
		log.Printf("[START] Minio Connection Check Error: %v", minioErr)
	}
	
	checkErr := database.CheckConnection()

	if checkErr != nil {
		log.Printf("[START] Databae Connection Check Error")
	}

	serving := openServer()

	log.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	log.Printf("[DEBUG] App Host %s", configs.GlobalConfig.AppHost)
	log.Printf("[START] Server Listening On: %s", configs.GlobalConfig.AppPort)
	log.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

	serving.ListenAndServe()
}

func openServer() *http.Server{
	server := http.NewServeMux()

	middlewares.CorsMiddlewares(server)
	routers.DefaultRouter(server)
	routers.AdminRouter(server)

	serving := &http.Server{
		Handler: 		server,
		Addr: 			configs.GlobalConfig.AppHost,
		WriteTimeout: 	30 * time.Second,
		ReadTimeout:  	30 * time.Second,
	}

	return serving
}

func setConfigs() {
	configs.SetGlobalConfig()
	configs.SetDatabaseConfig()
	configs.SetMinioConfig()
}
