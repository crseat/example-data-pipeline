package app

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"

	"github.com/crseat/example-data-pipeline/internal/domain"
)

type ConsumerService struct {
	reader     *kafka.Reader
	repository domain.Repository
}

func NewConsumerService(reader *kafka.Reader, repository domain.Repository) *ConsumerService {
	return &ConsumerService{reader: reader, repository: repository}
}

func (s *ConsumerService) ConsumeMessages() {
	for {
		message, err := s.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error reading message: %v", err)
			continue
		}

		var postData domain.PostData
		if err := json.Unmarshal(message.Value, &postData); err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue
		}

		if err := s.repository.SavePostData(postData); err != nil {
			log.Printf("error saving post data: %v", err)
		}
	}
}
