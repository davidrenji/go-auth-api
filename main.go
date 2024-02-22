package main

import (
	"log"
	"os"

	"github.com/davidrenji/go-bootcamp-api/connections"
	"github.com/davidrenji/go-bootcamp-api/routes"
	"github.com/davidrenji/go-bootcamp-api/utils"
	"github.com/gofiber/fiber/v2"
)

func init() {
	// LoadEnv()
	utils.LoadEnv()
	// InitDB()
	connections.InitDB()
}

func main() {
	// Create a new Fiber instance
	app := fiber.New()
	// Define routes
	routes.UserRoutes(app)
	// Start server on port xxxx
	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Error starting the server: %s", err)
	}
}
