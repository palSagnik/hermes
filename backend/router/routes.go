package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/palSagnik/hermes/handler"
	"github.com/palSagnik/hermes/middleware"
)

func SetUpRoutes(app *fiber.App) {

	app.Get("/alive", handler.Alive)

	// these routes would be for authentication purposes, hence grouped under auth
	auth := app.Group("/auth")
	auth.Post("/signup", handler.Signup)
	auth.Post("/login", handler.Login)
	auth.Get("/verify", handler.Verify)

	// these routes will only be accessible after verification
	// hence a token is used to access the api calls such as userdetails, expenses
	api := app.Group("/api", middleware.VerifyToken())
	api.Get("/users", handler.GetUsers)
}
