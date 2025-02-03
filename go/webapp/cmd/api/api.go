package main

import (
	"database/sql"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/maslias/webapp/cmd/configs"
	"github.com/maslias/webapp/internal/auth"
	"github.com/maslias/webapp/internal/store"
	"github.com/maslias/webapp/middleware"
)

type app struct {
	cfg        *configs.Config
	storage    *store.Storage
	db         *sql.DB
	logger     *zap.SugaredLogger
	auth       *auth.JWTAuthenticator
	middleware *middleware.Middleware
}

func NewApp(
	cfg *configs.Config,
	storage *store.Storage,
	db *sql.DB,
	logger *zap.SugaredLogger,
	auth *auth.JWTAuthenticator,
	midddleware *middleware.Middleware,
) *app {
	return &app{
		cfg:        cfg,
		storage:    storage,
		db:         db,
		logger:     logger,
		auth:       auth,
		middleware: midddleware,
	}
}

var UserInvitationExp time.Duration = time.Hour * 24 * 3

func (a *app) Run() error {
	router := http.NewServeMux()

	routerV1 := http.NewServeMux()
	router.Handle("/v1/", http.StripPrefix("/v1", routerV1))

	healthHandler := NewHealthHanddler(a.cfg.AppConfig, a.storage, a.db, a.logger)
	healthHandler.RegisterRoutes(routerV1)

	postsHandler := NewPostsHandler(a.storage, a.db, a.logger, a.middleware)
	postsHandler.registerRoutes(routerV1)

	commentsHandler := NewCommentsHandler(a.storage, a.db, a.logger)
	commentsHandler.RegisterRoutes(routerV1)

	authHandler := NewAuthHandler(a.storage, a.db, a.logger, a.cfg.AuthConfig, a.auth)
	authHandler.RegisterRoutes(routerV1)

	usersHandler := NewUsersHandler(a.storage, a.db, a.logger)
	usersHandler.RegisterRoutes(routerV1)

	server := http.Server{
		Addr:    a.cfg.APP_ADDR,
		Handler: a.middleware.Logger(router),
	}

	a.logger.Infow("server has started", "addr", a.cfg.APP_ADDR, "env", a.cfg.APP_ENV)
	return server.ListenAndServe()
}
