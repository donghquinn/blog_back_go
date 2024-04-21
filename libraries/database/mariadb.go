package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/donghquinn/blog_back_go/configs"
	"github.com/donghquinn/blog_back_go/queries"
	_ "github.com/go-sql-driver/mysql"
)

// DB 연결 인스턴스
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

// 테이블 생성
func CheckConnection() error {
	connect, dbErr := InitDatabaseConnection()

	if dbErr != nil {
		return dbErr
	}

	time.Sleep(time.Second * 5)
	
	pingErr := connect.Ping()

	if pingErr != nil {
		log.Printf("[DATABASE] Database Ping Error: %v", pingErr)

		return  pingErr
	}

	_, createUserErr := connect.Query(queries.CreateUserTable)

	if createUserErr != nil {
		log.Printf("[DATABASE] Create User Table Err: %v", createUserErr)

		return createUserErr
	}

	_, createPostErr := connect.Query(queries.CreatePostTable)

	if createPostErr != nil {
		log.Printf("[DATABASE] Create Post Table Error: %v", createPostErr)

		return createPostErr
	}
	
	defer connect.Close()

	return nil
}

// 쿼리
func Query(connect *sql.DB, queryString string, args ...string) (*sql.Rows, error) {
	var arguments []interface{}

    for _, arg := range args {
        arguments = append(arguments, arg)
    }	
	
	result, err := connect.Query(queryString, arguments...)

	if err != nil {
		log.Printf("[QUERY] Query Error: %v\n", err)

		return nil, err
	}

	return result, nil
}

// 인서트 쿼리
func InsertQuery(connect *sql.DB, queryString string, args ...string) (int64, error) {
	var arguments []interface{}

    for _, arg := range args {
        arguments = append(arguments, arg)
    }	

	insertResult, insertErr := connect.Exec(queryString, arguments...)

	if insertErr != nil {
		log.Printf("[INSERT] Insert Query Err: %v", insertErr)

		return -99999, insertErr
	}

	// Insert ID
	insertId, insertIdErr := insertResult.LastInsertId()

	if insertIdErr != nil {
		log.Printf("[INSERT] Get Insert ID Error: %v", insertIdErr)

		return -999999, insertIdErr
	}

	return insertId, nil
}