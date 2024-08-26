package main

import (
	"context"
	"errors"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	KAFKA_PRODUCER = "KAFKA"
	QUEUE_PRODUCER = "QUEUE" // rabbit mq
)

type Producer interface {
	SendMessage([]byte) error
	Close() error
}

type ProducerQueueImpl struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewProducer(producerType string) (Producer, error) {
	var producer Producer
	var err error
	switch producerType {
	case QUEUE_PRODUCER:
		producer, err = newProducerQueueImpl()
	case KAFKA_PRODUCER:
	// TODO
	default:
		return nil, errors.New("Not implemented")
	}

	return producer, err
}

func newProducerQueueImpl() (Producer, error) {
	url := os.Getenv("QUEUE_URL")
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &ProducerQueueImpl{conn, ch}, nil
}

func (p *ProducerQueueImpl) SendMessage(msg []byte) error {
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

func (p *ProducerQueueImpl) Close() error {
	return p.conn.Close()
}

type ProducerKafkaImpl struct {
}
