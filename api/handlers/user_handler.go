package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/user"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// CreateUser is a function to create user data to database
// @Summary Create User
// @Description Create User
// @Tags Users
// @Accept json
// @Produce json
// @Param user body entities.CreateUserSchema true "Register user"
// @Success 200 {object} presenter.JsonResponse{data=presenter.User}
// @Failure 503 {object} presenter.JsonResponse
// @Router /user [post]
func CreateUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		if requestBody.Email == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(errors.New(
				"Please specify title and author",
			)))
		}

		result, err := service.CreateUser(&requestBody)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(presenter.UserSuccessResponse(result))
	}
}
