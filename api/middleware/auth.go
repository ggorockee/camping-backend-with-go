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
	if err.Error() == "missing or malformed JWT" {
		jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
	}

	jsonResponse := presenter.NewJsonResponse(false, "Invalid or expired JWT", nil)
	return c.Status(http.StatusUnauthorized).JSON(jsonResponse)
}
