package presenter

import "github.com/gofiber/fiber/v2"

func AuthErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func AuthSuccessfulResponse(data any) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  false,
	}
}
