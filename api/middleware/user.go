package middleware

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func RoleMiddleware(allowedRoles ...entities.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Locals("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		userId := int(claims["user_id"].(float64))
		db := c.Locals("db").(*gorm.DB)
		var user entities.User
		if err := db.First(&user, userId).Error; err != nil {
			jsonResponse := presenter.JsonResponse{
				Error:   true,
				Message: "Access denied",
				Data:    nil,
			}
			return c.Status(fiber.StatusForbidden).JSON(jsonResponse)
		}

		// 사용자의 role이 허용된 role 목록에 있는지 확인
		for _, role := range allowedRoles {
			if user.Role == string(role) {
				return c.Next()
			}
		}

		jsonResponse := presenter.JsonResponse{
			Error:   true,
			Message: "Access denied",
			Data:    nil,
		}
		return c.Status(fiber.StatusForbidden).JSON(jsonResponse)
	}
}
