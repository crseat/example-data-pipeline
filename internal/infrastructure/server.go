package infrastructure

import (
	"github.com/crseat/example-data-pipeline/internal/adapters/http"
	"github.com/crseat/example-data-pipeline/internal/adapters/kafka"
	"github.com/crseat/example-data-pipeline/internal/app"
)

func StartServer() {
	// Load configuration
	config := LoadConfig()

	e := echo.New()

	// Initialize Kafka producer
	producer := kafka.NewKafkaProducer(config.KafkaBrokers, config.KafkaTopic)
	defer producer.Close()

	// Initialize service
	service := app.NewPostService(producer)

	// Initialize handler and register routes
	handler := http.NewHandler(service)
	handler.RegisterRoutes(e)

	// Start the server
	e.Logger.Fatal(e.Start(config.ServerPort))
}
