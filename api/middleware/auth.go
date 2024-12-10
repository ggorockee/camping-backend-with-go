package middleware

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/config"
	"camping-backend-with-go/pkg/entities"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
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

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Locals("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		userId := int(claims["user_id"].(float64))
		db := c.Locals("db").(*gorm.DB)
		var user entities.User
		if err := db.First(&user, userId).Error; err != nil {
			c.Locals("is_authenticated", false)
			c.Locals("request_user", "anonymous")
			return c.Next()
		}

		c.Locals("is_authenticated", true)
		c.Locals("request_user", user)
		return c.Next()
	}
}
