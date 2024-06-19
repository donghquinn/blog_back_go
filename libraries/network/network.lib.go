package network

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

func OpenServer() *http.Server{
	server := http.NewServeMux()

	middleWareHandler := middlewares.CorsMiddlewares(server)

	routers.DefaultRouter(server)
	routers.AdminRouter(server)

	serving := &http.Server{
		Handler: 		middleWareHandler,
		Addr: 			configs.GlobalConfig.AppHost,
		WriteTimeout: 	30 * time.Second,
		ReadTimeout:  	30 * time.Second,
	}

	return serving
}

func DatabaseConnect() {
	minioErr := database.MinioConnect()

	if minioErr != nil {
		log.Printf("[START] Minio Connection Check Error: %v", minioErr)
	}
	
	checkErr := database.CheckConnection()

	if checkErr != nil {
		log.Printf("[START] Databae Connection Check Error: %v", checkErr)
	}

	_, redisErr := database.RedisInstance()

	if redisErr != nil {
		log.Printf("[START] Redis Connection Check Error: %v", redisErr)
	}

}

func SetConfigs() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Printf("[ENV] Load Env Error")
	}
	
	configs.SetGlobalConfig()
	configs.SetDatabaseConfig()
	configs.SetMinioConfig()
	configs.SetRedisConfig()
}
