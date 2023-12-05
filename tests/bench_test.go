package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"testing"
)

func BenchmarkName(b *testing.B) {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "my-topic",
		Balancer: &kafka.LeastBytes{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			err := w.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte("Key-A"),
					Value: []byte("Hello World!"),
				},
			)
			if err != nil {
				log.Fatal("failed to write messages:", err)
			}
		}()
	}
}
