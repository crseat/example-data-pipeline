package kafka

import (
	"github.com/segmentio/kafka-go"
)

func NewKafkaConsumer(brokers []string, topic string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})
}
