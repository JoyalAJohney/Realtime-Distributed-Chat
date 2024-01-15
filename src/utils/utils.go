package utils

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	gonanoid "github.com/matoous/go-nanoid/v2"

	"realtime-chat/src/models"
)

func SendErrorMessage(conn *websocket.Conn, errMsg string) {
	errMessage := models.ErrorMessage{
		Error:   true,
		Message: errMsg,
	}

	if err := conn.WriteJSON(errMessage); err != nil {
		log.Println(err)
	}
}

func GenerateUserID() string {
	id, _ := gonanoid.New()
	return id
}
