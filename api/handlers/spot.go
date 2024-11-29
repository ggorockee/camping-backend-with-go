package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/spot"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// AddSpot is a function to create spot data to database
// @Summary AddSpot
// @Description AddSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param user body entities.CreateSpotInputSchema true "Create Spot"
// @Success 200 {object} presenter.JsonResponse{data=presenter.Spot}
// @Failure 503 {object} presenter.JsonResponse
// @Router /spot [post]
// @Security Bearer
func AddSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.CreateSpotInputSchema
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.SpotErrorResponse(err))
		}

		result, err := service.InsertSpot(&requestBody, c)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.SpotErrorResponse(err))
		}
		return c.JSON(presenter.SpotSuccessResponse(result))
	}
}

// GetSpots is a function to get all spot data from database
// @Summary GetSpots
// @Description GetSpots
// @Tags Spot
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{data=[]presenter.Spot}
// @Failure 503 {object} presenter.JsonResponse
// @Router /spot [get]
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

// UpdateSpot is a function to update spot data to database
// @Summary UpdateSpot
// @Description UpdateSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param id path int true "Spot id"
// @Param user body entities.UpdateSpotSchema true "Update Spot"
// @Success 200 {object} presenter.JsonResponse{data=presenter.Spot}
// @Failure 503 {object} presenter.JsonResponse
// @Router /spot/{id} [put]
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

// GetSpot is a function to get spot data to database
// @Summary GetSpot
// @Description GetSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param id path int true "Spot id"
// @Success 200 {object} presenter.JsonResponse{data=presenter.Spot}
// @Failure 503 {object} presenter.JsonResponse
// @Router /spot/{id} [get]
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

// PartialUpdateSpot is a function to update spot data to database
// @Summary PartialUpdateSpot
// @Description PartialUpdateSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param id path int true "Spot id"
// @Param user body entities.UpdateSpotSchema true "Update Spot"
// @Success 200 {object} presenter.JsonResponse{data=presenter.Spot}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/{id} [patch]
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

// RemoveSpot is a function to delete spot data to database
// @Summary RemoveSpot
// @Description RemoveSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param id path int true "Spot id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/{id} [delete]
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
