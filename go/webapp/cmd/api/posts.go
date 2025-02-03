package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/maslias/webapp/cmd/customerror"
	"github.com/maslias/webapp/cmd/utils"
	"github.com/maslias/webapp/internal/store"
	"github.com/maslias/webapp/middleware"
)

type PostsHandler struct {
	storage    *store.Storage
	db         *sql.DB
	logger     *zap.SugaredLogger
	middleware *middleware.Middleware
}

func NewPostsHandler(
	storage *store.Storage,
	db *sql.DB,
	logger *zap.SugaredLogger,
	middleware *middleware.Middleware,
) *PostsHandler {
	return &PostsHandler{
		storage:    storage,
		db:         db,
		logger:     logger,
		middleware: middleware,
	}
}

func (h *PostsHandler) registerRoutes(router *http.ServeMux) {
	routerPosts := http.NewServeMux()
	router.Handle(
		"/posts/",
		http.StripPrefix("/posts", h.middleware.AuthToken(routerPosts)),
	)

	routerPosts.HandleFunc("POST /", customerror.ErrorBridge(h.handleCreatePost, h.logger))
	routerPosts.HandleFunc("GET /{postId}", customerror.ErrorBridge(h.handleGetPostById, h.logger))
}

type payloadCreatePost struct {
	Content string   `json:"content" validate:"required,max=1000"`
	Title   string   `json:"title"   validate:"required,max=100"`
	Tags    []string `json:"tags"`
}

func (h *PostsHandler) handleGetPostById(w http.ResponseWriter, r *http.Request) error {
	postId := r.PathValue("postId")
	postIdIntVal, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		return customerror.NewHybridError(customerror.ErrInternalFailure, err)
	}

	ctx := r.Context()

	post, err := h.storage.Posts.GetById(ctx, postIdIntVal)
	if err != nil {
		switch {
		case errors.Is(err, customerror.ErrNotFound):
			return customerror.NewHybridError(customerror.ErrNotFound, err)
		default:
			return customerror.NewHybridError(customerror.ErrInternalFailure, err)
		}
	}

	comments, err := h.storage.Comments.GetByPostId(ctx, postIdIntVal)
	if err != nil {
		return customerror.NewHybridError(customerror.ErrInternalFailure, err)
	}

	post.Comments = comments

	return utils.WriteJSON(w, http.StatusFound, post)
}

func (h *PostsHandler) handleCreatePost(w http.ResponseWriter, r *http.Request) error {
	var payload *payloadCreatePost
	if err := utils.ReadLightJSON(w, r, &payload); err != nil {
		return customerror.NewHybridError(customerror.ErrBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		return customerror.NewHybridError(customerror.ErrBadRequest, err)
	}

	newPost := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserId:  1,
	}
	ctx := r.Context()

	if err := h.storage.Posts.Create(ctx, newPost); err != nil {
		return customerror.NewHybridError(customerror.ErrInternalFailure, err)
	}
	return utils.WriteJSON(w, http.StatusCreated, newPost)
}
