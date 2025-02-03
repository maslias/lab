package main

import (
	"context"
	"testing"

	"github.com/maslias/common/api"

	"github.com/maslias/microservices-payments/processor/inmen"
)

func TestService(t *testing.T) {
	processor := inmen.NewInmem()
	svc := NewService(processor)

	t.Run("should create paymentlink", func(t *testing.T) {
		link, err := svc.CreatePayment(context.Background(), &api.Order{})
		if err != nil {
			t.Errorf("createpayment() error = %v, want nil", err)
		}

		if link == "" {
			t.Error("link is empty")
		}
	})
}
