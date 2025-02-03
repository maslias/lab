//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/maslias/chatroomer/pkg/common"
	"github.com/maslias/chatroomer/pkg/config"
	"github.com/maslias/chatroomer/pkg/forwarder"
	"github.com/maslias/chatroomer/pkg/health"
	"github.com/maslias/chatroomer/pkg/user"
	"github.com/maslias/chatroomer/pkg/web"
)

func InitializeWeb() (*web.HttpServer, error) {
	wire.Build(
		config.NewConfig,
		common.NewHttpLog,
		web.NewHttpServer,
	)

	return &web.HttpServer{}, nil
}

func InitializeUser() (*user.GrpcServer, error) {
	wire.Build(
		config.NewConfig,
		common.NewHttpLog,
		user.NewGrpcServer,
	)
	return &user.GrpcServer{}, nil
}

func InitializeForwarder() (*forwarder.GrpcServer, error) {
	wire.Build(
		config.NewConfig,
		common.NewHttpLog,
		forwarder.NewGrpcServer,
	)
	return &forwarder.GrpcServer{}, nil
}

func InitliazeHealth() (*health.HttpServer, error) {
	wire.Build(
		config.NewConfig,
		common.NewHttpLog,
		health.NewHttpServer,
	)
	return &health.HttpServer{}, nil
}
