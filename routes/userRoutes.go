package routes

import (
	"github.com/davidrenji/go-bootcamp-api/controllers"
	"github.com/davidrenji/go-bootcamp-api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	// User Routes
	app.Route("/users", func(r fiber.Router) {
		r.Post("/", controllers.CreateUser)
		r.Post("/login", controllers.Login)
		r.Get("/", middlewares.ValidateAuth, controllers.GetUsers)
		r.Get("/:id", middlewares.ValidateAuth, controllers.GetUser)
		r.Put("/:id", middlewares.ValidateAuth, controllers.UpdateUser)
		r.Delete("/:id", middlewares.ValidateAuth, controllers.DeleteUser)
	})
}
