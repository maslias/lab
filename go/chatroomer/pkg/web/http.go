package web

import (
	"net/http"

	"github.com/maslias/chatroomer/pkg/common"
	"github.com/maslias/chatroomer/pkg/config"
)

type HttpServer struct {
	logger     *common.HttpLog
	name       string
	port       string
	version    string
	enviroment string
	server     *http.Server
}

func NewHttpServer(cfg *config.Config, logger *common.HttpLog) *HttpServer {
	return &HttpServer{
		port:       cfg.Web.Http.Server.Port,
		name:       cfg.Web.Http.Server.Name,
		version:    cfg.App.Version,
		enviroment: cfg.App.Enviroment,
		logger:     logger,
	}
}

func (srv *HttpServer) Run() error {
	router := http.NewServeMux()

	srv.server = &http.Server{
		Handler: router,
		Addr:    srv.port,
	}

	srv.logger.Infow(
		"http server is running",
		"name",
		srv.name,
		"port",
		srv.port,
		"env",
		srv.enviroment,
		"version",
		srv.version,
	)

	return srv.server.ListenAndServe()
}
