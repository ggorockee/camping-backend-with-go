package healthcheckhandler

import (
	"camping-backend-with-go/internal/domain/presenter"
	healthcheckservice "camping-backend-with-go/internal/domain/service/healthcheck"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetHealthCheck
// @Summary Health Check
// @Description Health Check
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /healthcheck [get]
func GetHealthCheck(service healthcheckservice.HealthCheckService) fiber.Handler {
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
