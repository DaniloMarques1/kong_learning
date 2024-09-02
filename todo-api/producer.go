package main

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	kafka "github.com/segmentio/kafka-go"
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
		producer, err = newProducerKafkaImpl()
	default:
		return nil, errors.New("Not implemented")
	}

	return producer, err
}

func newProducerQueueImpl() (Producer, error) {
	p := &ProducerQueueImpl{}
	url := p.getQueueUrl()
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	p.conn = conn
	p.ch = ch

	return p, nil
}

func (p *ProducerQueueImpl) getQueueUrl() string {
	url := os.Getenv("QUEUE_URL")
	if len(url) == 0 {
		return "amqp://fitz:fitz@localhost:5672"
	}
	return url
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
	conn *kafka.Conn
}

func newProducerKafkaImpl() (*ProducerKafkaImpl, error) {
	p := &ProducerKafkaImpl{}
	topic := "rank-topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", p.getConnectionString(), topic, partition)
	if err != nil {
		return nil, err
	}
	p.conn = conn
	return p, nil
}

func (kp *ProducerKafkaImpl) getConnectionString() string {
	host := os.Getenv("KAFKA_CONNECTION_STRING")
	if len(host) == 0 {
		return "localhost:9092"
	}
	return host
}

func (kp *ProducerKafkaImpl) SendMessage(msg []byte) error {
	kafkaMsg := kafka.Message{
		Key:       []byte(uuid.NewString()),
		Value:     msg,
		Partition: 0,
	}

	if _, err := kp.conn.WriteMessages(kafkaMsg); err != nil {
		return err
	}
	return nil
}

func (kp *ProducerKafkaImpl) Close() error {
	return kp.conn.Close()
}
