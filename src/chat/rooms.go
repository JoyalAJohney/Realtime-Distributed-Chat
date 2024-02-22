package chat

import (
	"context"
	"log"
	"os"
	"sync"

	"net/http"
	"encoding/json"
	"bytes"
	"strings"

	"realtime-chat/src/cache"
	"realtime-chat/src/kafka"
	"realtime-chat/src/models"
	"realtime-chat/src/utils"
)

var roomMutex = &sync.Mutex{}
var subscribedRooms = make(map[string]bool)

func JoinRoom(room string, user *models.User) {
	key := "room:" + room
	if err := addUserToRoomInRedis(key, user); err != nil {
		log.Println("Failed to add user to room:", err)
		utils.SendErrorMessage(user.Connection, "Unable to join room")
		return
	}

	cache.SubscribeToRoom(room, func(room string, message *models.Message) {
		BroadcastToRoom(room, *message)
	})

	message := models.Message{
		Sender:     user.ID,
		SenderName: user.Name,
		Room:       room,
		Type:       "join_room",
		Server:     os.Getenv("SERVER_NAME"),
	}
	cache.PublishMessage(room, &message)

	log.Printf("User %s joined room %s\n", user.ID, room)
	response := map[string]interface{}{
		"type":    "join_room",
		"room":    room,
		"success": true,
	}
	if err := user.Connection.WriteJSON(response); err != nil {
		log.Println("Error sending join room response:", err)
	}
}

func BroadcastToRoom(room string, message models.Message) {
	key := "room:" + room
	for _, userID := range getAllMembersInRoom(key) {
		// Get the websocket connection for the user from the local map
		if conn, exists := GetConnection(userID); exists {
			// Send the message to the user
			if err := conn.WriteJSON(message); err != nil {
				log.Printf("Error sending message to user %s: %v\n", userID, err)
				conn.Close()
				RemoveConnection(userID)
			}
		}
	}
}

func SendMessageToRoom(message models.Message, user *models.User) {
	cache.PublishMessage(message.Room, &message)
	kafka.PublishMessage(message)

	if strings.HasPrefix(message.Content, "@superchat") {
		response := GetResponseFromLLM(message.Content)
		log.Printf("Response from LLM")
		llmMessage := models.Message{
			Content: response,
			Sender: "SuperChat",
			SenderName: "SuperChat",
			Room: message.Room,
			Type: message.Type,
			Server: os.Getenv("SERVER_NAME"),
		}
		cache.PublishMessage(message.Room, &llmMessage)
	}
}


func GetResponseFromLLM(prompt string) string {
    requestBody := []byte(`
	{
		"model": "llama2", 
		"prompt": "` + prompt + `", 
		"stream": false, 
		"temperature": 0.8,
		"max_new_tokens": 75 
	}`)

    resp, err := http.Post("http://ollama:11434/api/generate", "application/json", bytes.NewBuffer(requestBody))
	
    if err != nil {
        log.Printf("Error calling AI service: %v", err)
        return "Error generating response"
    }

    defer resp.Body.Close()
    var response struct {
        Answer string `json:"response"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        log.Printf("Error decoding AI response: %v", err)
        return "Error processing response"
    }
	log.Printf("RAW Response from LLM: %+v", response)

    return response.Answer
}


func LeaveRoom(room string, user *models.User) {
	key := "room:" + room
	if err := removeUserFromRoomInRedis(key, user); err != nil {
		log.Println("Failed to remove user from room:", err)
		utils.SendErrorMessage(user.Connection, "Unable to leave room")
		return
	}

	cache.CheckAndUnsubscribeFromRoom(room)
	response := map[string]interface{}{
		"type":    "leave_room",
		"room":    room,
		"success": true,
	}
	if err := user.Connection.WriteJSON(response); err != nil {
		log.Println("Error sending leave room response:", err)
	}
}

func LeaveAllRooms(user *models.User) {
	for _, room := range cache.GetAllRooms() {
		isMember := isUserInRoom(room, user)
		if isMember {
			removeUserFromRoomInRedis(room, user)
		}
	}
}

// Helper methods
func addUserToRoomInRedis(room string, user *models.User) error {
	ctx := context.Background()
	_, err := cache.RedisClient.SAdd(ctx, room, user.ID).Result()
	return err
}

func removeUserFromRoomInRedis(room string, user *models.User) error {
	ctx := context.Background()
	_, err := cache.RedisClient.SRem(ctx, room, user.ID).Result()
	return err
}

func isUserInRoom(room string, user *models.User) bool {
	ctx := context.Background()
	isMember, err := cache.RedisClient.SIsMember(ctx, room, user.ID).Result()
	if err != nil {
		log.Println(err)
		return false
	}
	return isMember
}

func getAllMembersInRoom(room string) []string {
	ctx := context.Background()
	members, err := cache.RedisClient.SMembers(ctx, room).Result()
	if err != nil {
		log.Println(err)
		return nil
	}
	return members
}
