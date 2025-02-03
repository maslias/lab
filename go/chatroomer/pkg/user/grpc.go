package user

import (
	"net"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/maslias/chatroomer/pkg/common"
	"github.com/maslias/chatroomer/pkg/common/discovery"
	"github.com/maslias/chatroomer/pkg/config"
)

type GrpcServer struct {
	addr       string
	consulAddr string
	name       string
	server     *grpc.Server
	logger     *common.HttpLog
	version    string
	enviroment string
}

func NewGrpcServer(cfg *config.Config, logger *common.HttpLog) *GrpcServer {
	return &GrpcServer{
		addr:       cfg.User.Grpc.Server.Addr,
		consulAddr: cfg.Common.Consul.Addr,
		name:       cfg.User.Grpc.Server.Name,
		logger:     logger,
		version:    cfg.App.Version,
		enviroment: cfg.App.Enviroment,
	}
}

func (srv *GrpcServer) Run() error {
	consulRegistry, err := discovery.NewConsulRegistry(srv.consulAddr, srv.name)
	if err != nil {
		return err
	}

	ctx := context.Background()
	instanceId := discovery.GenerateConsulInstanceId(srv.name)

	if err := consulRegistry.Register(ctx, instanceId, srv.name, srv.addr); err != nil {
		return err
	}

	go func() {
		for {
			if err := consulRegistry.HealthCheck(instanceId, srv.name); err != nil {
				srv.logger.Errorw("error - userserver crashed", "error", err)
				os.Exit(1)
			}
			time.Sleep(time.Second * 1)

		}
	}()

	defer consulRegistry.Deregister(instanceId, srv.name)

	srv.server = grpc.NewServer()
	NewGrpcHandler(srv.server)

	ls, err := net.Listen("tcp", srv.addr)
	if err != nil {
		return err
	}
	defer ls.Close()

	srv.logger.Infow(
		"grpc server is running",
		"name",
		srv.name,
		"addr",
		srv.addr,
		"env",
		srv.enviroment,
		"version",
		srv.version,
	)

	err = srv.server.Serve(ls)
	if err != nil {
		return err
	}

	return nil
}
