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

// UpdateSpot is handler/controller which updates data of Spots in the Camping
func UpdateSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Spot

		err := c.BodyParser(&requestBody)
		id, _ := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.SpotErrorResponse(err))
		}
		result, err := service.UpdateSpot(&requestBody, id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.SpotErrorResponse(err))
		}
		return c.JSON(presenter.SpotSuccessResponse(result))
	}
}

// GetSpot is handler/controller which updates data of Spots in the Camping
func GetSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		fetched, err := service.GetSpot(id)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.SpotErrorResponse(err))
		}

		return c.JSON(presenter.SpotSuccessResponse(fetched))
	}
}

// PartialUpdateSpot is handler/controller which updates data of Spots in the Camping
func PartialUpdateSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Spot

		err := c.BodyParser(&requestBody)
		id, _ := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.SpotErrorResponse(err))
		}
		result, err := service.PartialUpdateSpot(&requestBody, id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.SpotErrorResponse(err))
		}
		return c.JSON(presenter.SpotSuccessResponse(result))
	}
}

// RemoveSpot is handler/controller which removes Books from the BookShop
func RemoveSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		err := service.RemoveSpot(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.SpotErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "delete successfully",
			"err":    nil,
		})
	}
}
