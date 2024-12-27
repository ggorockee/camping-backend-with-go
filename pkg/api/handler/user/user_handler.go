package userhandler

import (
	"camping-backend-with-go/internal/application/dto"
	"camping-backend-with-go/internal/domain/presenter"
	userservice "camping-backend-with-go/internal/domain/service/user"

	"github.com/gofiber/fiber/v2"
)

// ChangePassword
// @Summary ChangePassword
// @Description ChangePassword
// @Tags Users
// @Accept json
// @Produce json
// @Param requestBody body dto.ChangePasswordReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /user/change-password [put]
// @Security Bearer
func ChangePassword(service userservice.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.ChangePasswordReq

		// parsing error
		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		//oldPassword := requestBody.OldPassword
		newPassword := requestBody.NewPassword
		newPasswordConfirm := requestBody.NewPasswordConfirm

		// user가 존재하지 않음
		// service -> repository 에서 처리
		err := service.ChangePassword(&requestBody, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		// newPassword와 newPasswordConfirm이 다름
		// => 400 error
		if newPassword != newPasswordConfirm {
			jsonResponse := presenter.NewJsonResponse(true, "password didn't match", nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "success change password", nil)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}
