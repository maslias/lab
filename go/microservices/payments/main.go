package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/maslias/common"
	"github.com/maslias/common/broker"
	"github.com/maslias/common/discovery"
	"github.com/maslias/common/discovery/consul"
	"github.com/stripe/stripe-go/v80"
	"google.golang.org/grpc"

	stripeprocessor "github.com/maslias/microservices-payments/processor/stripe_processor"
)

var (
	grpcAddr              = common.GetString("GRPC_ADDR_PAYMENTS")
	consulAddr            = common.GetString("CONSUL_ADDR")
	serviceName           = "payments"
	amqpUser              = common.GetString("RABBITMQ_USER")
	amqpPass              = common.GetString("RABBITMQ_PASS")
	amqpHost              = common.GetString("RABBITMQ_HOST")
	amqpPort              = common.GetString("RABBITMQ_PORT")
	stripeKey             = common.GetString("STRIPE_KEY")
	stripeGatewayHttpAddr = common.GetString("STRIPE_GATEWAY_HTTP_ADDR")
)

func main() {
	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceId := discovery.GenerateInstanceId(serviceName)

	if err := registry.Register(ctx, instanceId, serviceName, grpcAddr); err != nil {
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

	stripe.Key = stripeKey

	brokerChan, brokerClose := broker.Connect(amqpUser, amqpPass, amqpHost, amqpPort)
	defer func() {
		brokerClose()
		brokerChan.Close()
	}()

	stripeProcessor := stripeprocessor.NewProcessor(stripeGatewayHttpAddr)
	svc := NewService(stripeProcessor)
	amqpConsumer := NewConsumer(svc)
	go amqpConsumer.Listen(brokerChan)

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer l.Close()

	log.Println("GRPC Server orders started at ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
