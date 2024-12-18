package authroute

import (
	authservice "camping-backend-with-go/internal/domain/service/auth"
	authhandler "camping-backend-with-go/pkg/api/handler/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service authservice.AuthService) {
	authRouter := app.Group("/auth")
	authRouter.Post("/signup", authhandler.CreateUser(service))
	authRouter.Post("/login", authhandler.Login(service))
}
