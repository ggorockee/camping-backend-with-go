package healthcheckroute

import (
	healthcheckservice "camping-backend-with-go/internal/domain/service/healthcheck"
	healthcheckhandler "camping-backend-with-go/pkg/api/handler/healthcheck"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckRouter(app fiber.Router, service healthcheckservice.HealthCheckService) {
	app.Get("/healthcheck", healthcheckhandler.GetHealthCheck(service))
}
