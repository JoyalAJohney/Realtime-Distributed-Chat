package processor

import (
	"log"
	"time"

	"kafka_consumer/database"
)

func ProcessMessages(messages []Message) {
	if len(messages) == 0 {
		return
	}

	for _, message := range messages {
		var currentTime = time.Now()
		dbMessage := database.DBMessage{
			UserID:    message.Sender,
			RoomID:    message.Room,
			Message:   message.Content,
			Timestamp: &currentTime,
		}
		log.Println("Saving message to database:", dbMessage)

		result := database.DB.CreateInBatches(&dbMessage, len(messages))
		if result.Error != nil {
			log.Printf("Error saving message to database: %v", result.Error)
		}
	}
}
