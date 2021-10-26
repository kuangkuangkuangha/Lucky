package config

func GetDbConfig() map[string]interface{} {

	// init db config
	dbConfig := make(map[string]interface{})

	dbConfig["hostname"] = "localhost"
	dbConfig["port"] = "3306"
	dbConfig["database"] = "lucky"
	dbConfig["username"] = "root"
	dbConfig["password"] = "zk2824895143"
	dbConfig["charset"] = "utf8mb4"
	dbConfig["parseTime"] = "True"

	return dbConfig
}
