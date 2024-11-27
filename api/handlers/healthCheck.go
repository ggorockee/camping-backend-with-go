package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/healthcheck"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// GetHealthCheck is a function to healthcheck
// @Summary Health Check
// @Description Health Check
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse
// @Failure 503 {object} presenter.JsonResponse
// @Router /healthcheck [get]
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
