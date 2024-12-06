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

		err := c.BodyParser(&requestBody)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		result, err := service.InsertSpot(&requestBody, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		userSerializer := entities.NewUserSerializer(&result.User)
		serializer := entities.NewSpotSerializer(result, userSerializer)
		jsonResponse := presenter.NewJsonResponse(true, err.Error(), serializer.DetailSerialize())

		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// GetMySpots is a function to get all spot data from database
// @Summary GetMySpots
// @Description GetMySpots
// @Tags Spot
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{data=[]entities.Spot}
// @Failure 503 {object} presenter.JsonResponse
// @Router /spot/me [get]
// @Security Bearer
func GetMySpots(service spot.Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchMySpots(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		spotsSerialized := make([]entities.SpotListOutputSchema, 0)

		for _, fetchedItem := range *fetched {
			userSerializer := entities.NewUserSerializer(&fetchedItem.User)
			serializer := entities.NewSpotSerializer(&fetchedItem, userSerializer)
			spotsSerialized = append(spotsSerialized, serializer.ListSerialize())
		}

		jsonResponse := presenter.NewJsonResponse(false, "", spotsSerialized)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
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
		var requestBody entities.UpdateSpotSchema

		err := c.BodyParser(&requestBody)
		id, _ := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusBadRequest).JSON(jsonResponse)
		}

		fetchedSpot, err := service.UpdateSpot(&requestBody, id, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		userSerializer := entities.NewUserSerializer(&fetchedSpot.User)
		serializer := entities.NewSpotSerializer(fetchedSpot, userSerializer)
		jsonResponse := presenter.NewJsonResponse(false, "", serializer.DetailSerialize())

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
// @Security Bearer
func GetSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, _ := c.ParamsInt("id")
		fetched, err := service.GetSpot(id, c)

		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		userSerializer := entities.NewUserSerializer(&fetched.User)
		serializer := entities.NewSpotSerializer(fetched, userSerializer)

		jsonResponse := presenter.NewJsonResponse(false, "", serializer.DetailSerialize())
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
// @Security Bearer
func RemoveSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, _ := c.ParamsInt("id")
		err := service.RemoveSpot(id, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "successfully delete", nil)
		return c.Status(http.StatusOK).JSON(jsonResponse)
	}
}

// GetAllSpots is a function to get all spot
// @Summary GetAllSpots
// @Description GetAllSpots
// @Tags Spot
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot [get]
func GetAllSpots(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		spots, err := service.GetAllSpots() //*[]entities.Spot

		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		responseSpots := make([]entities.SpotListOutputSchema, 0)
		for _, s := range *spots {
			userSerializer := entities.NewUserSerializer(&s.User)
			serializer := entities.NewSpotSerializer(&s, userSerializer)
			responseSpots = append(responseSpots, serializer.ListSerialize())
		}

		jsonResponse := presenter.NewJsonResponse(false, "", responseSpots)
		return c.Status(http.StatusOK).JSON(jsonResponse)
	}
}
