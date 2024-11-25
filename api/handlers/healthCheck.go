package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/healthcheck"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetHealthCheck(service healthcheck.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := service.GetHealthCheck()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.HealthCheckErrorResponse(err))
		}

		return c.Status(http.StatusOK).JSON(presenter.HealthCheckSuccessResponse())
	}
}
