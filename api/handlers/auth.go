package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/user"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Login is a function to Login
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body entities.LoginInputSchema true "Login Schema"
// @Success 200 {object} presenter.JsonResponse{data=string}
// @Failure 503 {object} presenter.JsonResponse
// @Router /auth/login [post]
func Login(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// request parser
		var requestBody entities.LoginInputSchema
		var jsonResponse presenter.JsonResponse
		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		token, err := service.Login(&requestBody)
		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		jsonResponse = presenter.JsonResponse{
			Error:   false,
			Message: "",
			Data:    token,
		}

		return c.Status(http.StatusOK).JSON(jsonResponse)
	}
}
