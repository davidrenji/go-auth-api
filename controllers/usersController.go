package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/davidrenji/go-bootcamp-api/connections"
	"github.com/davidrenji/go-bootcamp-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	// Get email/passs of request body
	body := models.User{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to bind request body"})
	}

	// Find user
	var user models.User
	result := connections.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid email or password")
	}

	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid email or password")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Invalid email or password")
	}

	// Return token
	return c.SendString(tokenString)
}

func CreateUser(c *fiber.Ctx) error {
	// Get the body from the request
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	// Encrypt the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to encrypt password")
	}

	// Set the password to the hashed password
	user.Password = string(hash)

	result := connections.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to create user")
	}

	return c.JSON(user)

}

func GetUsers(c *fiber.Ctx) error {

	var users []models.User

	result := connections.DB.Find(&users)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to get users")
	}

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {

	id := c.Params("id")

	var user models.User

	result := connections.DB.First(&user, id)

	if result.Error != nil {
		return c.Status(http.StatusNotFound).SendString("User not found")
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	result := connections.DB.First(&user, id)

	if result.Error != nil {
		return c.Status(http.StatusNotFound).SendString("User not found")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	// Encrypt the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to encrypt password")
	}

	// Set the password to the hashed password
	user.Password = string(hash)

	result = connections.DB.Save(&user)

	if result.Error != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to update user")
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	result := connections.DB.First(&user, id)

	if result.Error != nil {
		return c.Status(http.StatusNotFound).SendString("User not found")
	}

	result = connections.DB.Delete(&user)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to delete user")
	}

	return c.JSON(user)
}
