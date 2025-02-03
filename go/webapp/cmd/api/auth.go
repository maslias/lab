package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/maslias/webapp/cmd/configs"
	"github.com/maslias/webapp/cmd/customerror"
	"github.com/maslias/webapp/cmd/utils"
	"github.com/maslias/webapp/internal/auth"
	"github.com/maslias/webapp/internal/store"
)

type AuthHandler struct {
	storage *store.Storage
	db      *sql.DB
	logger  *zap.SugaredLogger
	cfg     *configs.AuthConfig
	auth    *auth.JWTAuthenticator
}

func NewAuthHandler(
	storage *store.Storage,
	db *sql.DB,
	logger *zap.SugaredLogger,
	cfg *configs.AuthConfig,
	auth *auth.JWTAuthenticator,
) *AuthHandler {
	return &AuthHandler{
		storage: storage,
		db:      db,
		logger:  logger,
		cfg:     cfg,
		auth:    auth,
	}
}

type PayloadCreateUser struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email"    validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type PayloadCreateToken struct {
	Email    string `json:"email"    validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (h *AuthHandler) RegisterRoutes(router *http.ServeMux) {
	routerAuth := http.NewServeMux()
	router.Handle("/authentication/", http.StripPrefix("/authentication", routerAuth))

	routerAuth.HandleFunc("/user/", customerror.ErrorBridge(h.handleUserRegistration, h.logger))
	routerAuth.HandleFunc("/token/", customerror.ErrorBridge(h.handleTokenRegistration, h.logger))
}

func (h *AuthHandler) handleTokenRegistration(w http.ResponseWriter, r *http.Request) error {
	// create token
	// insert token to db
	// response token

	var payload PayloadCreateToken
	if err := utils.ReadLightJSON(w, r, &payload); err != nil {
		return customerror.NewHybridError(customerror.ErrBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		return customerror.NewHybridError(customerror.ErrBadRequest, err)
	}

	ctx := r.Context()
	user, err := h.storage.Users.GetByEmail(ctx, payload.Email)
	if err != nil {
		return customerror.NewHybridError(customerror.ErrBadRequest, err)
	}

	claims := jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(h.cfg.AUTH_TOKEN_EXP).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": h.cfg.AUTH_TOKEN_ISS,
		"aud": h.cfg.AUTH_TOKEN_ISS,
	}

	token, err := h.auth.GenerateToken(claims)
	if err != nil {
		return customerror.NewHybridError(customerror.ErrInternalFailure, err)
	}

	return utils.WriteJSON(w, http.StatusCreated, token)
}

func (h *AuthHandler) handleUserRegistration(w http.ResponseWriter, r *http.Request) error {
	var payload PayloadCreateUser

	if err := utils.ReadLightJSON(w, r, &payload); err != nil {
		return customerror.NewHybridError(customerror.ErrBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		return customerror.NewHybridError(customerror.ErrBadRequest, err)
	}

	newUser := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
		RoleId:   1,
	}

	if err := newUser.Password.Set(payload.Password); err != nil {
		return customerror.NewHybridError(customerror.ErrInternalFailure, err)
	}

	ctx := r.Context()
	token := uuid.New().String()
	hash := sha256.Sum256([]byte(token))
	hashToken := hex.EncodeToString(hash[:])

	if err := h.storage.Users.CreateAndInvite(ctx, newUser, hashToken, UserInvitationExp); err != nil {
		return customerror.NewHybridError(customerror.ErrInternalFailure, err)
	}

	activationUrl := fmt.Sprintf("http://localhost:8080/v1/users/activate/{%s}", token)

	return utils.WriteJSON(w, http.StatusCreated, activationUrl)
}
