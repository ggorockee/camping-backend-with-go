package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/api/middleware"
	"camping-backend-with-go/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Patch("/user/changepw", middleware.Protected(), handlers.ChangePassword(service))
	app.Post("/user/changepw", middleware.Protected(), handlers.ChangePassword(service))
}
