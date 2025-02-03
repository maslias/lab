package handlers

import (
	"log/slog"
	"net/http"

)

type HTTPErrorHandler func(w http.ResponseWriter, r *http.Request) error

func ErrorBridge(h HTTPErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error(err.Error(), "method", r.Method, "path", r.URL.Path)
		}
	}
}

