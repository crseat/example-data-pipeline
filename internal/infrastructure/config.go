package infrastructure

import (
	"os"
	"strings"
)

type Config struct {
	ServerPort     string
	KafkaBrokers   []string
	KafkaTopic     string
	AppEnvironment string
	AerospikeHost  string
	AerospikePort  string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort:     getEnv("SERVER_PORT", ":8080"),
		KafkaBrokers:   strings.Split(getEnv("KAFKA_BROKER", "kafka:9092"), ","),
		KafkaTopic:     getEnv("KAFKA_TOPIC", "post-topic"),
		AppEnvironment: getEnv("APP_ENV", "development"),
		AerospikeHost:  getEnv("AEROSPIKE_HOST", "aerospike"),
		AerospikePort:  getEnv("AEROSPIKE_PORT", "3000"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
