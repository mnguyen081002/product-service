package infrastructure

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"productservice/internal/constants"
	"syscall"
)

type KafkaSubscribe struct {
	*kafka.Reader
}

func NewKafkaSubscribe(topic string) KafkaSubscribe {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		GroupID:   constants.GroupCmsProduct,
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	// wait signal to close
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		err := r.Close()
		if err != nil {
			fmt.Println("Error close kafka reader", err)
		}
		os.Exit(1)
	}()

	return KafkaSubscribe{
		r,
	}
}
