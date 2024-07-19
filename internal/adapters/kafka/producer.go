package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"

	"github.com/crseat/example-data-pipeline/internal/domain"
)

type Producer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *Producer) Produce(postData domain.PostData) error {
	message, err := json.Marshal(postData)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	})
}

func (p *Producer) Close() {
	if err := p.writer.Close(); err != nil {
		log.Printf("failed to close writer: %v", err)
	}
}
