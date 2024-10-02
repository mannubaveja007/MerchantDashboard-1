package config

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var KafkaProducer *kafka.Producer
func InitKafka(broker string) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	KafkaProducer = producer
}
func CloseKafka() {
	if KafkaProducer != nil {
		KafkaProducer.Close()
	}
}
