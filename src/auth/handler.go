package auth

import (
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"realtime-chat/src/config"
	"realtime-chat/src/database"
	"realtime-chat/src/models"
)

func SignUp(ctx *fiber.Ctx) error {
	var request models.AuthRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Println("Error parsing request body for Signup:", err)
		return ctx.JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request format. Please check your data.",
		})
	}

	if request.Username == "" || request.Password == "" {
		log.Println("Username or password is empty")
		return ctx.JSON(fiber.Map{
			"error":   true,
			"message": "Username and password are required fields.",
		})
	}

	var userExists database.DBUser
	if err := database.DB.Where("name = ?", request.Username).First(&userExists).Error; err == nil {
		log.Println("User already exists")
		return ctx.JSON(fiber.Map{
			"error":   true,
			"message": "Username already exists. Please choose a different one.",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return ctx.JSON(fiber.Map{
			"error":   true,
			"message": "An error occurred while processing your request. Please try again later.",
		})
	}

	user := database.DBUser{
		Name:     request.Username,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Error adding user to database: %v", err)
		return ctx.JSON(fiber.Map{
			"error":   true,
			"message": "An error occurred while processing your request. Please try again later.",
		})
	}

	return ctx.JSON(fiber.Map{
		"error":   false,
		"message": "Signup success",
	})
}

func Login(ctx *fiber.Ctx) error {
	var request models.AuthRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Printf("Error parsing request body for Login: %v", err)
		return ctx.JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request format. Please check your data.",
		})
	}

	if request.Username == "" || request.Password == "" {
		log.Println("Username or password is empty")
		return ctx.JSON(fiber.Map{
			"error":   true,
			"message": "Username and password are required fields.",
		})
	}

	var user database.DBUser
	if err := database.DB.Where("name = ?", request.Username).First(&user).Error; err != nil {
		log.Printf("Error finding user in database: %v", err)
		return ctx.JSON(fiber.Map{
			"error":   true,
			"message": "User not found. Please check your credentials.",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		log.Printf("Password mismatch: %v", err)
		return ctx.JSON(fiber.Map{
			"error":   true,
			"message": "Invalid password. Please try again.",
		})
	}

	userID := strconv.FormatUint(uint64(user.ID), 10)
	token, err := GenerateJWT(userID, user.Name)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return ctx.JSON(fiber.Map{
			"error":   true,
			"message": "An error occurred while processing your request. Please try again later.",
		})
	}

	return ctx.JSON(fiber.Map{
		"error":   false,
		"message": "Login success",
		"token":   token,
	})
}

func GenerateJWT(userID string, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte(config.Config.JwtSecret))
}
