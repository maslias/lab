package main

import (
	"context"
	"log"
	"net/http"
	"time"

	common "github.com/maslias/common"
	"github.com/maslias/common/discovery"
	"github.com/maslias/common/discovery/consul"
	"github.com/maslias/microservices-gateway/gateway"
)

var (
	appAddr     = common.GetString("APP_ADDR")
	consulAddr  = common.GetString("CONSUL_ADDR")
	serviceName = "gateway"
)

func main() {
	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceId := discovery.GenerateInstanceId(serviceName)

	if err := registry.Register(ctx, instanceId, serviceName, appAddr); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.HealthCheck(instanceId, serviceName); err != nil {
				log.Fatal("failed to heath check")
			}
			time.Sleep(time.Second * 1)
		}
	}()

	defer registry.Deregister(ctx, instanceId, serviceName)

	router := http.NewServeMux()
	apiRouter := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", apiRouter))

	ordersGateway := gateway.NewGrpcGateway(registry)

	handler := NewHandler(ordersGateway)
	handler.RegisterRoutes(apiRouter)

	log.Printf("start server at %s \n", appAddr)

	if err := http.ListenAndServe(appAddr, router); err != nil {
		log.Fatalf("failed to start server")
	}
}
