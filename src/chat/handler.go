package chat

import (
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"

	"realtime-chat/src/models"
)

func WebSocketHandler(conn *websocket.Conn) {
	userID := conn.Locals("userID").(string)
	userName := conn.Locals("userName").(string)

	user := &models.User{
		ID:         userID,
		Name:       userName,
		Connection: conn,
	}
	server := os.Getenv("SERVER_NAME")
	log.Printf("User %s connected to server: %s\n", userName, server)

	// Add connection for local tracking
	AddConnection(userID, conn)

	for {
		var message models.Message
		if err := conn.ReadJSON(&message); err != nil {
			log.Println("Error reading message from websocket:", err)
			break
		}

		log.Printf("Received message: %v\n on server: %s", message, server)
		message.Sender = userID
		message.Server = server
		message.SenderName = userName
		switch MessageType(message.Type) {
		case JoinRoomType:
			JoinRoom(message.Room, user)
		case LeaveRoomType:
			LeaveRoom(message.Room, user)
		case ChatMessageType:
			SendMessageToRoom(message, user)
		default:
			log.Println("Unknown message type")
		}
	}

	// User disconnected
	LeaveAllRooms(user)
	// Remove connection from local tracking
	RemoveConnection(userID)
}
