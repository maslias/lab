package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type ConfigData struct {
	APP_ADDR string
}

var Envs = initConfig()

func initConfig() ConfigData {
	err := godotenv.Load(".dev.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return ConfigData{
		APP_ADDR: getEnv("APP_ADDR"),
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal(fmt.Errorf("could not find env varliable: %v \n", key))
	}

	return value
}

func getEnvAsInt(key string) int64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal(fmt.Errorf("could not find env varliable: %v \n", key))
	}

	valueAsInt, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return valueAsInt
}
