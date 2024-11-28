package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/user"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
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

// CreateUser is a function to create user data to database
// @Summary Create User
// @Description Create User
// @Tags Users
// @Accept json
// @Produce json
// @Param user body entities.SignUpInputSchema true "Register user"
// @Success 200 {object} presenter.JsonResponse{data=presenter.User}
// @Failure 503 {object} presenter.JsonResponse
// @Router /user [post]
func CreateUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.SignUpInputSchema
		err := c.BodyParser(&requestBody)

		// json parsing
		if err != nil {
			jsonResponse := presenter.JsonResponse{
				Status: false,
				Data:   nil,
				Error:  err.Error(),
			}
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}
		if requestBody.Email == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(errors.New(
				"Please specify title and author",
			)))
		}

		err = service.CreateUser(&requestBody)
		if err != nil {
			jsonResponse := presenter.JsonResponse{
				Status: false,
				Data:   nil,
				Error:  err.Error(),
			}
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		// password와 confirm_password가 다르면 error
		if requestBody.Password != requestBody.PasswordConfirm {
			jsonResponse := presenter.JsonResponse{
				Status: false,
				Data:   nil,
				Error:  "password didn't match",
			}
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		jsonResponse := presenter.JsonResponse{
			Status: false,
			Data:   nil,
			Error:  "Welcome",
		}
		return c.Status(http.StatusOK).JSON(jsonResponse)
	}
}
