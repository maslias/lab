package health

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/maslias/chatroomer/pkg/common"
	"github.com/maslias/chatroomer/pkg/common/discovery"
	"github.com/maslias/chatroomer/pkg/config"
)

type HttpServer struct {
	logger     *common.HttpLog
	name       string
	port       string
	version    string
	enviroment string
	server     *http.Server
	consulAddr string
}

func NewHttpServer(cfg *config.Config, logger *common.HttpLog) *HttpServer {
	return &HttpServer{
		logger:     logger,
		name:       cfg.Health.Http.Server.Name,
		port:       cfg.Health.Http.Server.Port,
		version:    cfg.App.Version,
		enviroment: cfg.App.Enviroment,
		consulAddr: cfg.Common.Consul.Addr,
	}
}

func (srv *HttpServer) Run() error {
	consulRegistry, err := discovery.NewConsulRegistry(srv.consulAddr, srv.name)
	if err != nil {
		return err
	}

	ctx := context.Background()
	instanceId := discovery.GenerateConsulInstanceId(srv.name)

	if err := consulRegistry.Register(ctx, instanceId, srv.name, srv.port); err != nil {
		return err
	}

	go func() {
		for {
			if err := consulRegistry.HealthCheck(instanceId, srv.name); err != nil {
				srv.logger.Errorw("error - forwarder crashed", "error", err)
				os.Exit(1)
			}
			time.Sleep(time.Second * 1)

		}
	}()

	defer consulRegistry.Deregister(instanceId, srv.name)
	router := http.NewServeMux()

	routerHealth := http.NewServeMux()
	router.Handle("/health/", http.StripPrefix("/health", routerHealth))

	handlerHealth := NewHttpHandler(srv.logger)
	handlerHealth.RegisterRoutes(routerHealth)

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

	return srv.server.ListenAndServe()
}
