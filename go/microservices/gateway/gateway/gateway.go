package gateway

import (
	"context"

	pb "github.com/maslias/common/api"
)

type OrdersGateway interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error)
}
