package route

import (
	"camping-backend-with-go/internal_backup/handler"
	"camping-backend-with-go/pkg/service/user"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service user.Service) {
	app.Post("/auth/login", handler.Login(service))
}
