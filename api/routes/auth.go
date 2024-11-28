package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service user.Service) {
	app.Post("/auth/login", handlers.Login(service))
	app.Post("/auth/signup", handlers.CreateUser(service))
}
