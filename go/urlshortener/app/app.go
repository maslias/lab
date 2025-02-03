package app

import (
	"log/slog"
	"net/http"

	root "github.com/maslias/urlshortener"
	"github.com/maslias/urlshortener/handlers"
	"github.com/maslias/urlshortener/middleware"
)

type App struct {
	addr string
}

func NewApp(addr string) *App {
	return &App{addr: addr}
}

func (a *App) Run() error {
	router := http.NewServeMux()
    router.Handle("/public/", root.Public())


	homeHandler := handlers.NewIndexHandler()
	homeHandler.RegisterRoutes(router)
    

	server := http.Server{
		Addr:    a.addr,
		Handler: middleware.MiddlewareChain(router),
	}

    slog.Info("HTTP Server run", "listenAddr", a.addr)
	return server.ListenAndServe()
}
