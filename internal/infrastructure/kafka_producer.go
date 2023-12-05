package infrastructure

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Producer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type kafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer() Producer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	return &kafkaProducer{
		writer: w,
	}
}

func (k *kafkaProducer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	return k.writer.WriteMessages(ctx, msgs...)
}

func (k *kafkaProducer) Close() error {
	return k.writer.Close()
}
