package cache

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"realtime-chat/src/models"
)

var roomsMutex = &sync.Mutex{}
var subscribedRooms = make(map[string]bool)

var messageHandlerCallback models.MessageHandlerCallbackType

func SubscribeToRoom(room string, callback models.MessageHandlerCallbackType) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	// Subscribe to room if not already subscribed
	if !subscribedRooms[room] {
		ctx := context.Background()
		PubSubConnection.Subscribe(ctx, room)
		subscribedRooms[room] = true

		// Assign callback function to be used in listenForMessages()
		messageHandlerCallback = callback
		// Single goroutine to listen for messages on all subscribed channels
		go listenForMessages()
	}
}

// Listen for messages on the single PubSub connection
func listenForMessages() {
	channel := PubSubConnection.Channel()
	for message := range channel {
		log.Printf("Received message from channel: %s\n", message.Payload)
		var chatMessage models.Message
		err := json.Unmarshal([]byte(message.Payload), &chatMessage)
		if err != nil {
			log.Printf("Error decoding message from channel: %v\n", err)
			continue
		}

		// Call the callback function assigned in SubscribeToRoom()
		if messageHandlerCallback != nil {
			messageHandlerCallback(message.Channel, &chatMessage)
		}
	}
}

func CheckAndUnsubscribeFromRoom(room string) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	if subscribedRooms[room] {
		ctx := context.Background()
		key := "room:" + room
		members, _ := RedisClient.SCard(ctx, key).Result()

		// Unsubscribe from room if there are no members
		if members == 0 {
			PubSubConnection.Unsubscribe(ctx, room)
			delete(subscribedRooms, room)
		}
	}
}

func PublishMessage(room string, message *models.Message) {
	ctx := context.Background()
	msg, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return
	}
	// Any connection from the pool can be used to publish messages
	RedisClient.Publish(ctx, room, msg)
	log.Printf("Published message: %s to channel: %s\n", msg, room)
}

func GetAllRooms() []string {
	ctx := context.Background()
	keys, err := RedisClient.Keys(ctx, "room:*").Result()
	if err != nil {
		log.Println(err)
		return nil
	}
	return keys
}
