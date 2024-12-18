package userroute

import (
	userservice "camping-backend-with-go/internal/domain/service/user"
	userhandler "camping-backend-with-go/pkg/api/handler/user"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service userservice.UserService) {
	userRouter := app.Group("/user")
	userRouter.Put("/change-password", userhandler.ChangePassword(service))
}
