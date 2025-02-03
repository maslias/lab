package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/maslias/webapp/cmd/customerror"
	"github.com/maslias/webapp/cmd/utils"
	"github.com/maslias/webapp/internal/store"
)

type UsersHandler struct {
	storage *store.Storage
	db      *sql.DB
	logger  *zap.SugaredLogger
}

func NewUsersHandler(storage *store.Storage, db *sql.DB, logger *zap.SugaredLogger) *UsersHandler {
	return &UsersHandler{
		storage: storage,
		db:      db,
		logger:  logger,
	}
}

func (h *UsersHandler) RegisterRoutes(router *http.ServeMux) {
	routerUsers := http.NewServeMux()
	router.Handle("/users/", http.StripPrefix("/users", routerUsers))
	routerUsers.HandleFunc(
		"PUT /activate/{token}",
		customerror.ErrorBridge(h.handleActivate, h.logger),
	)
}

func (h *UsersHandler) handleActivate(w http.ResponseWriter, r *http.Request) error {
	token := r.PathValue("token")
	hash := sha256.Sum256([]byte(token))
	hashToken := hex.EncodeToString(hash[:])
    fmt.Printf("hashtoken %s\n", hashToken)
	ctx := r.Context()

	if err := h.storage.Users.ActivateFromInvitation(ctx, hashToken, UserInvitationExp); err != nil {
		return customerror.NewHybridError(customerror.ErrBadRequest, err)
	}

	return utils.WriteJSON(w, http.StatusAccepted, "user is activate")
}
