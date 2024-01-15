package main

import (
	"context"
	"encoding/json"
	"log"

	"kafka_consumer/database"
	"kafka_consumer/kafka"
	"kafka_consumer/processor"
)

var BATCH_SIZE = 3

func main() {
	database.InitPostgres()

	consumer := kafka.NewKafkaConsumer()
	defer consumer.Close()

	var batch []processor.Message

	for {
		m, err := consumer.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message from kafka:", err)
			continue
		}
		log.Println("Message received:", string(m.Value))

		var msg processor.Message
		if err := json.Unmarshal(m.Value, &msg); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		log.Println("Message unmarshalled:", msg)

		batch = append(batch, msg)
		if len(batch) >= BATCH_SIZE {
			log.Println("Processing batch:", batch)
			processor.ProcessMessages(batch)
			batch = []processor.Message{}
		}
	}

}
