package configs

import (
	"time"

	"github.com/maslias/webapp/internal/envs"
)

type AppConfig struct {
	APP_ADDR    string
	APP_VERSION string
	APP_ENV     string
}

type DbConfig struct {
	DB_DRIVER         string
	DB_ADDR           string
	DB_MAX_OPEN_CONNS int
	DB_MAX_IDLE_CONNS int
	DB_MAX_IDLE_TIME  time.Duration
	DB_CTX_TIMEOUT    time.Duration
}

type AuthConfig struct {
	AUTH_TOKEN_SECRET string
	AUTH_TOKEN_EXP    time.Duration
	AUTH_TOKEN_ISS    string
}

type Config struct {
	*AppConfig
	*DbConfig
	*AuthConfig
}

func NewConfig() *Config {
	return &Config{
		&AppConfig{
			APP_ADDR:    envs.GetString("APP_ADDR"),
			APP_VERSION: envs.GetString("APP_VERSION"),
			APP_ENV:     envs.GetString("APP_ENV"),
		},
		&DbConfig{
			DB_DRIVER:         envs.GetString("DB_DRIVER"),
			DB_ADDR:           envs.GetString("DB_ADDR"),
			DB_MAX_OPEN_CONNS: envs.GetInt("DB_MAX_OPEN_CONNS"),
			DB_MAX_IDLE_CONNS: envs.GetInt("DB_MAX_IDLE_CONNS"),
			DB_MAX_IDLE_TIME:  envs.GetTimeDuration("DB_MAX_IDLE_TIME"),
			DB_CTX_TIMEOUT:    envs.GetTimeDuration("DB_CTX_TIMEOUT"),
		},
		&AuthConfig{
			AUTH_TOKEN_SECRET: envs.GetString("AUTH_TOKEN_SECRET"),
			AUTH_TOKEN_EXP:    time.Hour * 24 * 3,
			AUTH_TOKEN_ISS:    "webapp",
		},
	}
}
