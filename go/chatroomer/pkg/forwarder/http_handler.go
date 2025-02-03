package forwarder

import (
	"errors"
	"net/http"

	"github.com/maslias/chatroomer/pkg/common"
)

type HttpHandler struct {
	logger *common.HttpLog
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (h *HttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /users", h.handleCreateUser)
}

func (h *HttpHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var payload PayloadCreateUser
	if err := common.ReadJSON(w, r, payload); err != nil {
		h.errorRespsone(w, r, errors.Join(common.ErrBadRequest, err))
		return
	}

	if err := common.Validate.Struct(payload); err != nil {
		h.errorRespsone(w, r, errors.Join(common.ErrBadRequest, err))
		return
	}


}

func (h *HttpHandler) errorRespsone(w http.ResponseWriter, r *http.Request, err error) {
	status := common.GetErrorHttpStatus(err)
	if err := common.WriteJSONError(w, status, err); err == nil {
		h.logger.Errorw(
			"could not errorResponse from http server",
			"error",
			err.Error(),
			"status",
			status,
			"name",
		)
		return
	}

	h.logger.Errorw(
		err.Error(),
		"status",
		status,
		"method",
		r.Method,
		"path",
		r.URL.Path,
	)
}
