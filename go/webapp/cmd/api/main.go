package main

import (
	"fmt"
	"log"

	"go.uber.org/zap"

	"github.com/maslias/webapp/cmd/configs"
	"github.com/maslias/webapp/internal/auth"
	"github.com/maslias/webapp/internal/db"
	"github.com/maslias/webapp/internal/store"
	"github.com/maslias/webapp/middleware"
)

func main() {
	cfg := configs.NewConfig()

    fmt.Println(cfg)
	dbconn, err := db.NewDb(
		cfg.DbConfig.DB_DRIVER,
		cfg.DbConfig.DB_ADDR,
		cfg.DbConfig.DB_MAX_OPEN_CONNS,
		cfg.DbConfig.DB_MAX_IDLE_CONNS,
		cfg.DbConfig.DB_MAX_IDLE_TIME,
		cfg.DbConfig.DB_CTX_TIMEOUT,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer dbconn.Close()

	jwtAuth := auth.NewJWTAuthenticator(
		cfg.AuthConfig.AUTH_TOKEN_SECRET,
		cfg.AuthConfig.AUTH_TOKEN_ISS,
		cfg.AuthConfig.AUTH_TOKEN_ISS,
	)

	storage := store.NewStorage(dbconn)

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	middleware := middleware.NewMiddleware(logger, jwtAuth, storage)

	app := NewApp(cfg, storage, dbconn, logger, jwtAuth, middleware)
	if err := app.Run(); err != nil {
		logger.Fatal(err)
	}
}
