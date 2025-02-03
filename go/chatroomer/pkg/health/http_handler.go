package health

import (
	"fmt"
	"net/http"

	"github.com/maslias/chatroomer/pkg/common"
)

type HttpHandler struct {
	logger *common.HttpLog
}

func NewHttpHandler(logger *common.HttpLog) *HttpHandler {
	return &HttpHandler{logger: logger}
}

func (h *HttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/health", h.handleHealth)
}

func (h *HttpHandler) handleHealth(w http.ResponseWriter, r *http.Request) {
	h.logger.Infow("hanlder is called", "service", "health", "handler", "check")
	fmt.Printf("hello from: %s : %s\n", "health", "check")
	return
}
