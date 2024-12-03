package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/pkg/service/healthcheck"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckRouter(app fiber.Router, service healthcheck.Service) {
	app.Get("/healthcheck", handlers.GetHealthCheck(service))
}
