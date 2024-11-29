package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/api/middleware"
	"camping-backend-with-go/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Put("/user/changepw", middleware.Protected(), handlers.ChangePassword(service))
}
