package authhandler

import (
	authdto "camping-backend-with-go/internal/application/dto/auth"
	"camping-backend-with-go/internal/domain/presenter"
	authservice "camping-backend-with-go/internal/domain/service/auth"

	"github.com/gofiber/fiber/v2"
)

// CreateUser
// @Summary Create User
// @Description Create User
// @Tags Auth
// @Accept json
// @Produce json
// @Param requestBody body authdto.SignUpReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /auth/signup [post]
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
			jsonResponse := presenter.NewJsonResponse(true, "password가 다릅니다.", nil)
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
// @Param requestBody body authdto.LoginReq true "requestBody"
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
