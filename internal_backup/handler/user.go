package handler

import (
	"camping-backend-with-go/internal_backup/presenter"
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/service/user"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ChangePassword is a function to ChangePassword
// @Summary ChangePassword
// @Description ChangePassword
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.ChangePasswordIn true "Change Password"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /user/changepw [put]
// @Security Bearer
func ChangePassword(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.ChangePasswordIn

		// parsing error
		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		//oldPassword := requestBody.OldPassword
		newPassword := requestBody.NewPassword
		newPasswordConfirm := requestBody.NewPasswordConfirm

		// user가 존재하지 않음
		// service -> repository 에서 처리
		err := service.ChangePassword(&requestBody, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		// newPassword와 newPasswordConfirm이 다름
		// => 400 error
		if newPassword != newPasswordConfirm {
			jsonResponse := presenter.NewJsonResponse(true, "password didn't match", nil)
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "success change password", nil)
		return c.Status(http.StatusOK).JSON(jsonResponse)
	}
}

// CreateUser is a function to create user data to database
// @Summary Create User
// @Description Create User
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.SignUpIn true "Register user"
// @Success 200 {object} presenter.JsonResponse{data=entities.User}
// @Failure 503 {object} presenter.JsonResponse
// @Router /user/signup [post]
func CreateUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.SignUpIn
		err := c.BodyParser(&requestBody)

		// json parsing
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		err = service.CreateUser(&requestBody)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		// password와 confirm_password가 다르면 error
		if requestBody.Password != requestBody.PasswordConfirm {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "welcome!!", nil)
		return c.Status(http.StatusOK).JSON(jsonResponse)
	}
}
