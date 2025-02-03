package common

import (
	"log"
	"os"
)

func GetString(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("could not find env with key %s \n", key)
	}
	return val
}
