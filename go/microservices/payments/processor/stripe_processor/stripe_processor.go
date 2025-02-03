package stripeprocessor

import (
	"fmt"
	"log"

	pb "github.com/maslias/common/api"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/checkout/session"
)

type Stripe struct {
	gatewayHttpAddr string
}

func NewProcessor(gatewayHttpAddr string) *Stripe {
	return &Stripe{
		gatewayHttpAddr,
	}
}

func (s *Stripe) CreatePaymentLink(o *pb.Order) (string, error) {
	log.Printf("creating paymentlink for order %v\n", o)

	gatewaySuccessUrl := fmt.Sprintf("%s/success.html", s.gatewayHttpAddr)

	items := []*stripe.CheckoutSessionLineItemParams{}

	for _, item := range o.Items {
		items = append(items, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String("price_1QASEOK0x5N8wHj4miQWv8w1"),
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}

	params := &stripe.CheckoutSessionParams{
		LineItems: items,
		// AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
		// Customer:     stripe.String("{{CUSTOMER_ID}}"),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(gatewaySuccessUrl),
		// CancelURL:    stripe.String("https://example.com/cancel"),
	}

	result, err := session.New(params)
	if err != nil {
		return "", err
	}

	return result.URL, nil
}
