package env

import (
	"os"

	env "github.com/joho/godotenv"
)

const port_key string = "PORT"
const db_key string = "DBFILE"
const cfg_path_key string = "CONFIG_PATH"

func getByKey(key string) string {
	err := env.Load("local.env")

	if err != nil {
		panic("Невозможно загрузить .ENV")
	}

	return os.Getenv(key)
}

func GetConfigPath() string {
	return getByKey(cfg_path_key)
}

func GetPort() string {
	return getByKey(port_key)
}

func GetDbName() string {
	return getByKey(db_key)
}
