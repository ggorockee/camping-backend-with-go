package route

import (
	"camping-backend-with-go/pkg/service/user"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Post("/user/signup", handler.CreateUser(service))
	//app.Put("/user/changepw",
	//	middleware.Protected(),
	//	handler.ChangePassword(service),
	//)
	app.Put("/user/changepw",
		handler.ChangePassword(service),
	)
}
