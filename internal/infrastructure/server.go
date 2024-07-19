package infrastructure

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"

	"github.com/crseat/example-data-pipeline/internal/adapters/http"
	"github.com/crseat/example-data-pipeline/internal/adapters/kafka"
	"github.com/crseat/example-data-pipeline/internal/adapters/repositories"
	"github.com/crseat/example-data-pipeline/internal/app"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func StartServer() {
	// Load configuration
	config := LoadConfig()

	// Initialize Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Initialize Kafka producer
	producer := kafka.NewKafkaProducer(config.KafkaBrokers, config.KafkaTopic)
	defer producer.Close()

	// Initialize Aerospike repository
	repository, err := repositories.NewAerospikeRepository(config.AerospikeHost, config.AerospikePort)
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
