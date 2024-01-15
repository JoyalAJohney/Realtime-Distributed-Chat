package main

import (
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"realtime-chat/src/auth"
	"realtime-chat/src/cache"
	"realtime-chat/src/chat"
	"realtime-chat/src/config"
	"realtime-chat/src/database"
)

func main() {
	app := fiber.New()
	cache.InitRedis()
	database.InitPostgres()

	// Handle CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "POST, OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Health check
	app.Get("/api/health", func(ctx *fiber.Ctx) error {
		server := os.Getenv("SERVER_NAME")
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "healthy", "server": server})
	})

	app.Post("/api/auth/signup", auth.SignUp)
	app.Post("/api/auth/login", auth.Login)

	// Secure websocket connection
	app.Use("/ws", upgradeToWebSocket)
	app.Get("/ws/chat", websocket.New(chat.WebSocketHandler))

	port := ":" + config.Config.ServerPort
	log.Fatal(app.Listen(port))
}

// Authorize and Upgrate to websocket
func upgradeToWebSocket(context *fiber.Ctx) error {
	token := context.Query("token")
	if token == "" {
		log.Println("No token provided")
		return fiber.ErrUnauthorized
	}

	// Validate JWT token
	if err := auth.ValidateJWTToken(token); err != nil {
		log.Println("Error validating JWT token:", err)
		return fiber.ErrUnauthorized
	}

	userID, userName := auth.ParseJWTToken(token)
	if websocket.IsWebSocketUpgrade(context) {
		context.Locals("allowed", true)
		context.Locals("userID", userID)
		context.Locals("userName", userName)
		return context.Next()
	}
	return fiber.ErrUpgradeRequired
}
