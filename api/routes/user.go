package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Post("/user", handlers.CreateUser(service))
}
