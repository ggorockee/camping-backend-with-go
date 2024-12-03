package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/api/middleware"
	"camping-backend-with-go/pkg/service/user"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Post("/user/signup", handlers.CreateUser(service))
	app.Put("/user/changepw", middleware.Protected(), handlers.ChangePassword(service))
}
