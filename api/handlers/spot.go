package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/spot"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// AddSpot is a function to create spot data to database
// @Summary AddSpot
// @Description AddSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param user body entities.CreateSpotInputSchema true "Create Spot"
// @Success 200 {object} presenter.JsonResponse{data=entities.Spot}
// @Failure 503 {object} presenter.JsonResponse
// @Router /spot [post]
// @Security Bearer
func AddSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.CreateSpotInputSchema
		var jsonResponse presenter.JsonResponse
		err := c.BodyParser(&requestBody)
		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		result, err := service.InsertSpot(&requestBody, c)
		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse = presenter.JsonResponse{
			Error:   false,
			Message: "successfully create spot!",
			Data:    result.ListSerialize(),
		}
		return c.JSON(jsonResponse)
	}
}

// GetSpots is a function to get all spot data from database
// @Summary GetSpots
// @Description GetSpots
// @Tags Spot
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{data=[]entities.Spot}
// @Failure 503 {object} presenter.JsonResponse
// @Router /spot [get]
// @Security Bearer
func GetSpots(service spot.Service) fiber.Handler {
	var jsonResponse presenter.JsonResponse
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchSpots()
		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse = presenter.JsonResponse{
			Error:   false,
			Message: "",
			Data:    fetched,
		}
		return c.JSON(jsonResponse)
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
// @Success 200 {object} presenter.JsonResponse{data=entities.Spot}
// @Failure 503 {object} presenter.JsonResponse
// @Router /spot/{id} [put]
// @Security Bearer
func UpdateSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Spot
		var jsonResponse presenter.JsonResponse

		err := c.BodyParser(&requestBody)
		id, _ := c.ParamsInt("id")
		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}
		result, err := service.UpdateSpot(&requestBody, id)
		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}

			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse = presenter.JsonResponse{
			Error:   false,
			Message: "",
			Data:    result,
		}
		return c.Status(http.StatusOK).JSON(jsonResponse)
	}
}

// GetSpot is a function to get spot data to database
// @Summary GetSpot
// @Description GetSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param id path int true "Spot id"
// @Success 200 {object} presenter.JsonResponse{data=entities.Spot}
// @Failure 503 {object} presenter.JsonResponse
// @Router /spot/{id} [get]
func GetSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var jsonResponse presenter.JsonResponse
		id, _ := c.ParamsInt("id")
		fetched, err := service.GetSpot(id, c)

		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}

			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse = presenter.JsonResponse{
			Error:   false,
			Message: "",
			Data:    fetched,
		}
		return c.Status(http.StatusOK).JSON(jsonResponse)
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
// @Success 200 {object} presenter.JsonResponse{data=entities.Spot}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/{id} [patch]
func PartialUpdateSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Spot
		var jsonResponse presenter.JsonResponse

		err := c.BodyParser(&requestBody)
		id, _ := c.ParamsInt("id")
		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}
		result, err := service.PartialUpdateSpot(&requestBody, id)
		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse = presenter.JsonResponse{
			Error:   false,
			Message: "",
			Data:    result,
		}
		return c.Status(http.StatusOK).JSON(jsonResponse)
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
		var jsonResponse presenter.JsonResponse
		id, _ := c.ParamsInt("id")
		err := service.RemoveSpot(id)
		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse = presenter.JsonResponse{
			Error:   false,
			Message: "delete successfully",
			Data:    nil,
		}
		return c.Status(http.StatusOK).JSON(jsonResponse)
	}
}
