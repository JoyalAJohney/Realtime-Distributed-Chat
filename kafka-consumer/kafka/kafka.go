package kafka

import (
	"github.com/segmentio/kafka-go"

	"kafka_consumer/config"
)

type KafkaConfig struct {
	Host       string
	Port       string
	kafkaTopic string
	GroupID    string
}

var kafkaConfig KafkaConfig

func init() {
	kafkaConfig = KafkaConfig{
		Host:       config.Config.KafkaHost,
		Port:       config.Config.KafkaPort,
		kafkaTopic: config.Config.KafkaTopic,
		GroupID:    config.Config.KafkaGroupID,
	}
}

func NewKafkaConsumer() *kafka.Reader {
	kafkaBrokerAddress := kafkaConfig.Host + ":" + kafkaConfig.Port
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBrokerAddress},
		Topic:   kafkaConfig.kafkaTopic,
		GroupID: kafkaConfig.GroupID,
	})
}
