package forwarder

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
		logger:     logger,
		name:       cfg.Forwarder.Http.Server.Name,
		port:       cfg.Forwarder.Http.Server.Port,
		version:    cfg.App.Version,
		enviroment: cfg.App.Enviroment,
	}
}

func (srv *HttpServer) Run() error {
	router := http.NewServeMux()

	srv.server = &http.Server{
		Addr:    srv.port,
		Handler: router,
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
	err := srv.server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
