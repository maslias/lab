package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/maslias/common/api"
	"github.com/maslias/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service    OrdersService
	brokerChan *amqp.Channel
}

func NewGrpcHandler(grpcServer *grpc.Server, service OrdersService, brokerChan *amqp.Channel) {
	handler := &grpcHandler{
		service:    service,
		brokerChan: brokerChan,
	}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(
	ctx context.Context,
	payload *pb.CreateOrderRequest,
) (*pb.Order, error) {
	log.Printf("new order received: %v \n", payload)

	// o := &pb.Order{
	// 	Id: "42",
	// }

	o, err := h.service.CreateOrder(ctx, payload)
	if err != nil {
		return nil, err
	}

	marshallOrder, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	q, err := h.brokerChan.QueueDeclare(broker.OrderCreatedEvent, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	h.brokerChan.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType:  "application/json",
		Body:         marshallOrder,
		DeliveryMode: amqp.Persistent,
	})

	return o, nil
}
