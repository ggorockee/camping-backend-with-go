package presenter

import "github.com/gofiber/fiber/v2"

func HealthCheckSuccessResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   "fiber ready!!",
		"error":  nil,
	}
}

func HealthCheckErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
