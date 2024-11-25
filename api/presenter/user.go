package presenter

import "github.com/gofiber/fiber/v2"

func JwtErrorResponse(data *[]Spot) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}
