package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/auth"
	"camping-backend-with-go/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Login(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// request parser
		var requestBody entities.Login
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(http.StatusBadRequest).JSON(presenter.AuthErrorResponse(err))
		}

		token, err := service.Login(&requestBody)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(presenter.AuthErrorResponse(err))
		}

		return c.Status(http.StatusOK).JSON(presenter.AuthSuccessfulResponse(token))
	}
}
