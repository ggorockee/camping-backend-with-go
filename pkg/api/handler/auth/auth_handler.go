package authhandler

import (
	authdto "camping-backend-with-go/internal/application/dto/auth"
	"camping-backend-with-go/internal/domain/presenter"
	authservice "camping-backend-with-go/internal/domain/service/auth"

	"github.com/gofiber/fiber/v2"
)

// ChangePassword
// @Summary ChangePassword
// @Description ChangePassword
// @Tags Users
// @Accept json
// @Produce json
// @Param requestBody body authdto.ChangePasswordReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /user/change-password [put]
// @Security Bearer
func ChangePassword(service authservice.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody authdto.ChangePasswordReq

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

// CreateUser
// @Summary Create User
// @Description Create User
// @Tags Users
// @Accept json
// @Produce json
// @Param requestBody body authdto.SignUpReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /user/signup [post]
func CreateUser(service authservice.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody authdto.SignUpReq
		err := c.BodyParser(&requestBody)

		// json parsing
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		err = service.CreateUser(&requestBody)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		// password와 confirm_password가 다르면 error
		if requestBody.Password != requestBody.PasswordConfirm {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "welcome!!", nil)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// Login
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param requestBody body dto.LoginIn true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse
// @Router /auth/login [post]
func Login(service authservice.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// request parser
		var requestBody authdto.LoginReq

		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		token, err := service.Login(&requestBody)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "", token)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}
