package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/davidrenji/go-bootcamp-api/connections"
	"github.com/davidrenji/go-bootcamp-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateAuth(c *fiber.Ctx) error {
	//To handle any error
	defer handlePanic(c)

	// Get token
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		panic("Invalid token")
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		panic("Invalid token")
	}

	// Validate token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			panic("Invalid token")
		}

		// Lets check the if user exists in the database
		var user models.User
		result := connections.DB.First(&user, int(claims["id"].(float64)))
		if result.Error != nil {
			panic("Invalid token")
		}

		c.Locals("user", user)
	} else {
		panic("Invalid token")
	}

	return nil
}

func handlePanic(c *fiber.Ctx) error {
	if r := recover(); r != nil {
		return c.Status(http.StatusUnauthorized).SendString("Invalid token")
	}
	//If no panic,
	return c.Next()
}
