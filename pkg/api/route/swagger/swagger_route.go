package swaggerroute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRouter(app fiber.Router) {
	app.Get("/docs/*", swagger.HandlerDefault)
}
