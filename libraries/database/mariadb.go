package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/donghquinn/blog_back_go/configs"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabaseConnection() (*sql.DB, error) {
	dbConfig := configs.DatabaseConfig

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
	dbConfig.User, 
	dbConfig.Password,
	dbConfig.Host, 
	dbConfig.Port, 
	dbConfig.Database)

	connect, err := sql.Open("mysql", dbUrl)

	if err != nil {
		log.Printf("[DATABASE] Start Database Connection Error: %v", err)

		return nil, err
	}

	connect.SetConnMaxLifetime(time.Second * 30)
	connect.SetMaxOpenConns(10)
	connect.SetMaxIdleConns(10)
	connect.SetMaxOpenConns(30)

	return connect, nil
}

func CheckConnection() error {
	connect, dbErr := InitDatabaseConnection()

	if dbErr != nil {
		return dbErr
	}
	
	pingErr := connect.Ping()

	if pingErr != nil {
		log.Printf("[DATABASE] Database Ping Error: %v", pingErr)

		return  pingErr
	}

	return nil
}