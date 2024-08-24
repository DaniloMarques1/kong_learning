package main

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer interface {
	SendMessage([]byte) error
	Close() error
}

type ProducerImpl struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewProducer(url string) (*ProducerImpl, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &ProducerImpl{conn, ch}, nil
}

func (p *ProducerImpl) SendMessage(msg []byte) error {
	queue, err := p.ch.QueueDeclare("scheduler-queue", true, false, false, false, nil)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = p.ch.PublishWithContext(
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{ContentType: "text/plain", Body: msg},
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProducerImpl) Close() error {
	return p.conn.Close()
}
