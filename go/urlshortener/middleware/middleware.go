package middleware

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type Middleware func(http.Handler) http.Handler

var MiddlewareChain = middlewareChainProgress(registerLogger)

func middlewareChainProgress(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
            m := middlewares[i]
            next = m(next)
		}
		return next
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

func registerLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)

		slog.Info(
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
