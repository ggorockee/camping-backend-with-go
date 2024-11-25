package presenter

import (
	"camping-backend-with-go/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func UserSuccessResponse(data *entities.User) *fiber.Map {
	user := User{
		Id:       data.Id,
		Email:    data.Email,
		Username: data.Username,
	}
	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}
