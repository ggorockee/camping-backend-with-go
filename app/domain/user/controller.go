package user

import (
	"camping-backend-with-go/app"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Table() []app.Mapping
}

type controller struct {
	service Service
}

func (ctrl *controller) Table() []app.Mapping {
	v1 := "/api/v1"
	return []app.Mapping{
		{
			Method:  fiber.MethodGet,
			Path:    v1 + "/hello",
			Handler: ctrl.HelloWorld,
		},
	}
}

// HelloWorld godoc
// @Summary Say hello
// @Description Returns a hello world message
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello, World!"
// @Router /hello [get]
func (u *controller) HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func NewController(s Service) Controller {
	return &controller{
		service: s,
	}
}
