package route

import (
	"camping-backend-with-go/internal_backup/handler"
	"camping-backend-with-go/pkg/service/healthcheck"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckRouter(app fiber.Router, service healthcheck.Service) {
	app.Get("/healthcheck", handler.GetHealthCheck(service))
}
