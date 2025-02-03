package config

import "os"


type ConfigData struct {
	DB_DRIVER string
	DB_PATH   string
}

var config ConfigData

func Load() {
	config = ConfigData{
        DB_PATH: os.Getenv("XDG_CACHE_HOME") + "/tasks.db",
		DB_DRIVER: "sqlite3",
	}
}

func GET() ConfigData {
    return config
}
