package models

import (
	"github.com/gofiber/contrib/websocket"
)

type User struct {
	ID         string
	Name       string
	Connection *websocket.Conn
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MessageHandlerCallbackType func(room string, message *Message)

type Message struct {
	Sender     string `json:"sender"`
	SenderName string `json:"senderName"`
	Room       string `json:"room"`
	Type       string `json:"type"`
	Content    string `json:"content"`
	Server     string `json:"server"`
}

type ErrorMessage struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
