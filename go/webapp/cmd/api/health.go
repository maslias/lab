package main

import (
	"database/sql"
	"net/http"

	"github.com/maslias/webapp/cmd/configs"
	"github.com/maslias/webapp/cmd/customerror"
	"github.com/maslias/webapp/cmd/utils"
	"github.com/maslias/webapp/internal/store"
	"go.uber.org/zap"
)

type HealthHandler struct {
	cfg     *configs.AppConfig
	storage *store.Storage
	db      *sql.DB
    logger *zap.SugaredLogger
}

func NewHealthHanddler(cfg *configs.AppConfig, storage *store.Storage, db *sql.DB, logger *zap.SugaredLogger) *HealthHandler {
	return &HealthHandler{
		cfg:     cfg,
		storage: storage,
		db:      db,
        logger: logger,
	}
}

func (h *HealthHandler) RegisterRoutes(router *http.ServeMux) {
	routerHealth := http.NewServeMux()
	router.Handle("/health/", http.StripPrefix("/health", routerHealth))
	routerHealth.HandleFunc("/", customerror.ErrorBridge(h.handleHealth, h.logger))
}

func (h *HealthHandler) handleHealth(w http.ResponseWriter, r *http.Request) error {
	data := map[string]string{
		"status":  "ok",
		"env":     h.cfg.APP_ENV,
		"version": h.cfg.APP_VERSION,
	}

	return utils.WriteJSON(w, http.StatusOK, data)
}
