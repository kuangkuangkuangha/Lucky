package config

import (
	"os"
)

var dbConfig map[string]interface{}

func init() {
	// init db config

	dbConfig = make(map[string]interface{})

	dbConfig["hostname"] = os.Getenv("db_server")
	dbConfig["port"] = os.Getenv("db_port")
	dbConfig["database"] = os.Getenv("db_name")
	dbConfig["username"] = os.Getenv("db_username")
	dbConfig["password"] = os.Getenv("db_password")

	dbConfig["charset"] = "utf8"
	dbConfig["parseTime"] = "True"
	dbConfig["timezone"] = "Asia%2fShanghai"
	dbConfig["maxIdleConns"] = 20
	dbConfig["maxOpenConns"] = 100

}

func GetDbConfig() map[string]interface{} {
	return dbConfig
}
