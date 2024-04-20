package main

import (
	"log"
	"net/http"
	"time"

	"github.com/donghquinn/blog_back_go/configs"
	"github.com/donghquinn/blog_back_go/libraries/database"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Printf("[ENV] Load Env Error")
	}

	configs.SetGlobalConfig()
	configs.SetDatabaseConfig()
	configs.SetMinioConfig()

	checkErr := database.CheckConnection()

	if checkErr != nil {
		log.Printf("[START] Databae Connection Check Error")
	}

	server := http.NewServeMux()

	serving := &http.Server {
		Handler: server,
		Addr: configs.GlobalConfig.AppHost,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	log.Printf("[START] Server Listening On: %s", configs.GlobalConfig.AppPort)
	log.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

	serving.ListenAndServe()
}
