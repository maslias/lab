package main

import (
	"context"
	"log"

	common "github.com/maslias/common"
	pb "github.com/maslias/common/api"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{
		store: store,
	}
}

func (s *service) CreateOrder(
	ctx context.Context,
	payload *pb.CreateOrderRequest,
) (*pb.Order, error) {
	items, err := s.ValidateOrder(ctx, payload)
	if err != nil {
		return nil, err
	}

	o := &pb.Order{
		Id:         "42",
		CustomerId: payload.CustomerId,
		Status:     "pending",
		Items:      items,
	}

	return o, nil
}

func (s *service) ValidateOrder(
	ctx context.Context,
	payload *pb.CreateOrderRequest,
) ([]*pb.Item, error) {
	if len(payload.Items) == 0 {
		return nil, common.ErrNoItems
	}

	mergedItems := mergeItemsQuantities(payload.Items)
	log.Printf("merged items: %v \n", mergedItems)

	// validate with stock service
	// temp:
	var itemsWithPrice []*pb.Item
	for _, i := range mergedItems {
		itemsWithPrice = append(itemsWithPrice, &pb.Item{
			PriceId:  "price_1QASEOK0x5N8wHj4miQWv8w1",
			Id:       i.Id,
			Quantity: i.Quantity,
		})
	}

	return itemsWithPrice, nil
}

func mergeItemsQuantities(items []*pb.ItemsWithQantity) []*pb.ItemsWithQantity {
	merged := make([]*pb.ItemsWithQantity, 0)

	for _, item := range items {
		found := false
		for _, finalItem := range merged {
			finalItem.Quantity += item.Quantity
			found = true
			break
		}
		if !found {
			merged = append(merged, item)
		}
	}

	return merged
}
