package envs

import (
	"log"
	"os"
	"strconv"
	"time"
)

func GetString(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("could not load env varliable with key %s \n", key)
	}
	return val
}

func GetInt(key string) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("could not load env varliable with key %s \n", key)
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
        log.Fatalf("strconv key %s: %s\n",key, err.Error())
	}

	return valInt
}

func GetTimeDuration(key string) time.Duration {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("could not load env varliable with key %s \n", key)
	}

	valTd, err := time.ParseDuration(val)
	if err != nil {
        log.Fatalf("parseduration key %s: %s\n",key, err.Error())
	}

	return valTd
}
