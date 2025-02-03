package main

import (
	"context"

	pb "github.com/maslias/common/api"

	"github.com/maslias/microservices-payments/processor"
)

type service struct {
	processor processor.PaymentProcessor
}

func NewService(processor processor.PaymentProcessor) *service {
	return &service{processor}
}

func (s *service) CreatePayment(ctx context.Context, o *pb.Order) (string, error) {
	link, err := s.processor.CreatePaymentLink(o)
	if err != nil {
		return "", err
	}

	return link, nil
}
