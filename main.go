package main

import (
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
	app := fiber.New()

	routes.UserRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
