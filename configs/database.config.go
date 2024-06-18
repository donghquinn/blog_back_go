package configs

import "os"


type DatabaseConfigStruct struct {
	Host string
	Port string
	Database string
	User string
	Password string
}

var DatabaseConfig DatabaseConfigStruct


func SetDatabaseConfig() {
	DatabaseConfig.Host = os.Getenv("MARIADB_HOST")
	DatabaseConfig.User = os.Getenv("MARIADB_USER")
	DatabaseConfig.Password = os.Getenv("MARIADB_PASSWORD")
	DatabaseConfig.Port = os.Getenv("MARIADB_PORT")
	DatabaseConfig.Database = os.Getenv("MARIADB_DATABASE")
}
