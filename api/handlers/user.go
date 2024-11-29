package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/user"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ChangePassword(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.ChangePasswordInputSchema
		var jsonResponse presenter.JsonResponse

		// parsing error
		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse = presenter.JsonResponse{
				Status: false,
				Data:   nil,
				Error:  err.Error(),
			}
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		//oldPassword := requestBody.OldPassword
		newPassword := requestBody.NewPassword
		newPasswordConfirm := requestBody.NewPasswordConfirm

		// user가 존재하지 않음
		// service -> repository 에서 처리
		err := service.ChangePassword(&requestBody, c)
		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Status: false,
				Data:   nil,
				Error:  err.Error(),
			}
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		// newPassword와 newPasswordConfirm이 다름
		// => 400 error
		if newPassword != newPasswordConfirm {
			jsonResponse = presenter.JsonResponse{
				Status: false,
				Data:   nil,
				Error:  "password didn't not match",
			}
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		jsonResponse = presenter.JsonResponse{
			Status: true,
			Data:   "success change password",
			Error:  "",
		}
		return c.Status(http.StatusOK).JSON(jsonResponse)
	}
}
