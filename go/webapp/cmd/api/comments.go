package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/maslias/webapp/cmd/customerror"
	"github.com/maslias/webapp/cmd/utils"
	"github.com/maslias/webapp/internal/store"
	"go.uber.org/zap"
)

type CommentsHandler struct {
	storage *store.Storage
	db      *sql.DB
    logger *zap.SugaredLogger
}

func NewCommentsHandler(storage *store.Storage, db *sql.DB, logger *zap.SugaredLogger) *CommentsHandler {
	return &CommentsHandler{
		storage: storage,
		db:      db,
        logger: logger,
	}
}

func (h *CommentsHandler) RegisterRoutes(router *http.ServeMux) {
	routerComments := http.NewServeMux()
	router.Handle("/comments/", http.StripPrefix("/comments", routerComments))

	routerComments.HandleFunc("POST /", customerror.ErrorBridge(h.handleCreateComment, h.logger))
	routerComments.HandleFunc("GET /{commentId}", customerror.ErrorBridge(h.handleGetCommentById, h.logger))
}

type payloadCreateComment struct {
	content string `json:"content" validate:"required,max=1000"`
}

func (h *CommentsHandler) handleGetCommentById(w http.ResponseWriter, r *http.Request) error {
	commentId := r.PathValue("commentId")
	commentIdIntVal, err := strconv.ParseInt(commentId, 10, 64)
	if err != nil {
		return customerror.NewHybridError(customerror.ErrInternalFailure, err)
	}

	ctx := r.Context()
	comment, err := h.storage.Comments.GetById(ctx, commentIdIntVal)
	if err != nil {
		switch {
		case errors.Is(err, customerror.ErrNotFound):
			return customerror.NewHybridError(customerror.ErrNotFound, err)
		default:
			return customerror.NewHybridError(customerror.ErrInternalFailure, err)
		}
	}

	return utils.WriteJSON(w, http.StatusFound, comment)
}

func (h *CommentsHandler) handleCreateComment(w http.ResponseWriter, r *http.Request) error {
	var payload *payloadCreateComment
	if err := utils.ReadLightJSON(w, r, &payload); err != nil {
		return customerror.NewHybridError(customerror.ErrBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		return customerror.NewHybridError(customerror.ErrBadRequest, err)
	}

	newComment := &store.Comment{
		UserId:  1,
		PostId:  1,
		Content: payload.content,
	}

	ctx := r.Context()
	if err := h.storage.Comments.Create(ctx, newComment); err != nil {
		return customerror.NewHybridError(customerror.ErrInternalFailure, err)
	}

	return utils.WriteJSON(w, http.StatusCreated, newComment)
}
