package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/maslias/common/api"
	"github.com/maslias/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	service PaymentService
}

func NewConsumer(service PaymentService) *consumer {
	return &consumer{
		service: service,
	}
}

func (c *consumer) Listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(broker.OrderCreatedEvent, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	var closeChan chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("received message: %s\n", d.Body)
			o := &pb.Order{}
			if err := json.Unmarshal(d.Body, o); err != nil {
				log.Printf("failed to unmarshal order: %v \n", err)
				continue
			}

			paymenLink, err := c.service.CreatePayment(context.Background(), o)
			if err != nil {
				log.Printf("failed to create paymen: %v\n", err)
				continue
			}

            log.Printf("paymentlink: %s \n", paymenLink)
		}
	}()

	<-closeChan
}
