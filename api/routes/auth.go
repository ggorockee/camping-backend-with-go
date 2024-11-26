package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service auth.Service) {
	app.Post("/auth/login", handlers.Login(service))
}
