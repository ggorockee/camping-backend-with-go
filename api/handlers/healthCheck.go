package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/service/healthcheck"
	"net/http"

	"github.com/gofiber/fiber/v2"
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
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "welcome", nil)
		return c.Status(http.StatusOK).JSON(jsonResponse)
	}
}
