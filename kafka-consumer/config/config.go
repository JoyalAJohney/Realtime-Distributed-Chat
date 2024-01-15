package config

import (
	"log"
	"os"
)

type AppConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	KafkaHost        string
	KafkaPort        string
	KafkaTopic       string
	KafkaGroupID     string
}

var Config AppConfig

func init() {
	// Initialize Configuration
	Config = AppConfig{
		KafkaHost:        os.Getenv("KAFKA_HOST"),
		KafkaPort:        os.Getenv("KAFKA_PORT"),
		KafkaTopic:       os.Getenv("KAFKA_TOPIC"),
		KafkaGroupID:     os.Getenv("KAFKA_GROUP_ID"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDatabase: os.Getenv("POSTGRES_DATABASE"),
	}

	if Config.KafkaHost == "" || Config.PostgresHost == "" {
		log.Fatal("Environment variables not set")
	}
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
