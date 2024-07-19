package app

import "github.com/crseat/example-data-pipeline/internal/domain"

type PostService struct {
	producer PostProducer
}

type PostProducer interface {
	Produce(postData domain.PostData) error
}

func NewPostService(producer PostProducer) *PostService {
	return &PostService{producer: producer}
}

func (s *PostService) ProcessPostData(postData domain.PostData) error {
	return s.producer.Produce(postData)
}
