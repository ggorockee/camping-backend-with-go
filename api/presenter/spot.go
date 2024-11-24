package presenter

import (
	"camping-backend-with-go/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type Spot struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func SpotSuccessResponse(data *entities.Spot) *fiber.Map {
	spot := Spot{
		Id:     data.Id,
		Title:  data.Title,
		Author: data.Author,
	}
	return &fiber.Map{
		"status": true,
		"data":   spot,
		"error":  nil,
	}
}

func SpotsSuccessResponse(data *[]Spot) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func SpotErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
