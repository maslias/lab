package gateway

import (
	"context"

	pb "github.com/maslias/common/api"
	"github.com/maslias/common/discovery"
)

type gateway struct {
	registry discovery.Registry
}

func NewGrpcGateway(registry discovery.Registry) *gateway {
	return &gateway{
		registry: registry,
	}
}

func (g *gateway) CreateOrder(
	ctx context.Context,
	payload *pb.CreateOrderRequest,
) (*pb.Order, error) {
	conn, err := discovery.ServiceConnection(ctx, "orders", g.registry)
	if err != nil {
		return nil, err
	}

	c := pb.NewOrderServiceClient(conn)
	return c.CreateOrder(ctx, &pb.CreateOrderRequest{
		CustomerId: payload.CustomerId,
		Items:      payload.Items,
	})
}
