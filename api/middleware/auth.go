package middleware

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/config"
	"net/http"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   jwtware.SigningKey{Key: []byte(config.Config("JWT_SECRET"))},
		AuthScheme:   "Bearer",
		TokenLookup:  "header:Authorization",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	var jsonResponse presenter.JsonResponse
	if err.Error() == "missing or malformed JWT" {
		jsonResponse = presenter.JsonResponse{
			Error:   true,
			Data:    nil,
			Message: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
	}

	jsonResponse = presenter.JsonResponse{
		Error:   true,
		Data:    nil,
		Message: "Invalid or expired JWT",
	}
	return c.Status(http.StatusUnauthorized).JSON(jsonResponse)
}
