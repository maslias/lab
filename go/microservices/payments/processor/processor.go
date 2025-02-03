package processor

import pb "github.com/maslias/common/api"

type PaymentProcessor interface {
	CreatePaymentLink(*pb.Order) (string, error)
}
