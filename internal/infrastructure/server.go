package infrastructure

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/crseat/example-data-pipeline/internal/adapters/aerospike"
	"github.com/crseat/example-data-pipeline/internal/adapters/http"
	"github.com/crseat/example-data-pipeline/internal/adapters/kafka"
	"github.com/crseat/example-data-pipeline/internal/app"
)

func StartServer() {
	// Load configuration
	config := LoadConfig()

	// Initialize Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize Kafka producer
	producer := kafka.NewKafkaProducer(config.KafkaBrokers, config.KafkaTopic)
	defer producer.Close()

	// Initialize Aerospike repository
	repository, err := aerospike.NewAerospikeRepository(config.AerospikeHost, config.AerospikePort)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize service
	service := app.NewPostService(producer)

	// Initialize handler and register routes
	handler := http.NewHandler(service)
	handler.RegisterRoutes(e)

	// Initialize Kafka consumer
	consumer := kafka.NewKafkaConsumer(config.KafkaBrokers, config.KafkaTopic, "example-consumer-group")
	consumerService := app.NewConsumerService(consumer, repository)

	// Start Kafka consumer in a separate goroutine
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		consumerService.ConsumeMessages()
	}()

	// Start the Echo server
	e.Logger.Fatal(e.Start(config.ServerPort))

	// Wait for Kafka consumer goroutine to finish
	wg.Wait()
}
