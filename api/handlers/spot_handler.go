package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/spot"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AddSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Spot
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.SpotErrorResponse(err))
		}
		if requestBody.Author == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.SpotErrorResponse(errors.New(
				"Please specify title and author",
			)))
		}

		result, err := service.InsertSpot(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.SpotErrorResponse(err))
		}
		return c.JSON(presenter.SpotSuccessResponse(result))
	}
}

func GetSpots(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchSpots()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.SpotErrorResponse(err))
		}
		return c.JSON(presenter.SpotsSuccessResponse(fetched))
	}
}
