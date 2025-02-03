package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"

	"github.com/maslias/webapp/cmd/customerror"
	"github.com/maslias/webapp/cmd/utils"
	"github.com/maslias/webapp/internal/auth"
	"github.com/maslias/webapp/internal/store"
)

type Middleware struct {
	logger  *zap.SugaredLogger
	auth    *auth.JWTAuthenticator
	storage *store.Storage
}

func NewMiddleware(
	logger *zap.SugaredLogger,
	auth *auth.JWTAuthenticator,
	storage *store.Storage,
) *Middleware {
	return &Middleware{
		logger:  logger,
		auth:    auth,
		storage: storage,
	}
}

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) writeHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (mw *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)

		mw.logger.Infow(
			strconv.Itoa(wrapped.statusCode),
			"method",
			r.Method,
			"path",
			r.URL.Path,
			"duration",
			time.Since(start),
		)
	})
}

func (mw *Middleware) AuthToken(
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := w.Header().Get("Authorization")
		if authHeader == "" {
			err := customerror.NewHybridError(
				customerror.ErroUnauth,
				fmt.Errorf("authorization header is missing"),
			)
			utils.WriteJSONError(w, http.StatusNonAuthoritativeInfo, err)
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			err := customerror.NewHybridError(
				customerror.ErroUnauth,
				fmt.Errorf("authorization header is malformed"),
			)
			utils.WriteJSONError(w, http.StatusNonAuthoritativeInfo, err)
			return
		}

		token := authParts[1]
		jwtToken, err := mw.auth.ValidateToken(token)
		if err != nil {
			err := customerror.NewHybridError(
				customerror.ErroUnauth,
				err,
			)

			utils.WriteJSONError(w, http.StatusNonAuthoritativeInfo, err)
			return
		}

		claims, _ := jwtToken.Claims.(jwt.MapClaims)
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
		if err != nil {
			err := customerror.NewHybridError(
				customerror.ErroUnauth,
				err,
			)

			utils.WriteJSONError(w, http.StatusNonAuthoritativeInfo, err)
			return
		}

		ctx := r.Context()

		user, err := mw.storage.Users.GetById(ctx, userId)
		if err != nil {
			err := customerror.NewHybridError(
				customerror.ErroUnauth,
				err,
			)
			utils.WriteJSONError(w, http.StatusNonAuthoritativeInfo, err)
            return 
		}

		ctx = context.WithValue(ctx, "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
