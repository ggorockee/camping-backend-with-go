package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/serializer"
	"camping-backend-with-go/pkg/service/spot"
	"gorm.io/gorm"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// AddSpot is a function to create spot data to database
// @Summary AddSpot
// @Description AddSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param user body dto.CreateSpotIn true "Create Spot"
// @Success 200 {object} presenter.JsonResponse{data=entities.Spot}
// @Failure 503 {object} presenter.JsonResponse
// @Router /spot [post]
// @Security Bearer
func AddSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.CreateSpotIn
		db := c.Locals("db").(*gorm.DB)

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

		userSerializer := serializer.NewUserSerializer(&result.User)
		categorySerializer := serializer.NewCategorySerializer(&result.Category)
		spotSerializer := serializer.NewSpotSerializer(result, userSerializer, categorySerializer)
		jsonResponse := presenter.NewJsonResponse(true, "", spotSerializer.DetailSerialize(db))

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
		db := c.Locals("db").(*gorm.DB)
		fetched, err := service.FetchMySpots(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		spotsSerialized := make([]dto.SpotListOut, 0)

		for _, fetchedItem := range *fetched {
			userSerializer := serializer.NewUserSerializer(&fetchedItem.User)
			categorySerializer := serializer.NewCategorySerializer(&fetchedItem.Category)
			spotSerializer := serializer.NewSpotSerializer(&fetchedItem, userSerializer, categorySerializer)
			spotsSerialized = append(spotsSerialized, spotSerializer.ListSerialize(db))
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
// @Param user body dto.UpdateSpotIn true "Update Spot"
// @Success 200 {object} presenter.JsonResponse{data=entities.Spot}
// @Failure 503 {object} presenter.JsonResponse
// @Router /spot/{id} [put]
// @Security Bearer
func UpdateSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		var requestBody dto.UpdateSpotIn

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

		userSerializer := serializer.NewUserSerializer(&fetchedSpot.User)
		categorySerializer := serializer.NewCategorySerializer(&fetchedSpot.Category)
		spotSerializer := serializer.NewSpotSerializer(fetchedSpot, userSerializer, categorySerializer)
		jsonResponse := presenter.NewJsonResponse(false, "", spotSerializer.DetailSerialize(db))

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
		db := c.Locals("db").(*gorm.DB)

		id, _ := c.ParamsInt("id")
		fetched, err := service.GetSpot(id, c)

		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		userSerializer := serializer.NewUserSerializer(&fetched.User)
		categorySerializer := serializer.NewCategorySerializer(&fetched.Category)
		spotSerializer := serializer.NewSpotSerializer(fetched, userSerializer, categorySerializer)

		jsonResponse := presenter.NewJsonResponse(false, "", spotSerializer.DetailSerialize(db, c))
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
		db := c.Locals("db").(*gorm.DB)
		spots, err := service.GetAllSpots() //*[]entities.Spot

		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(http.StatusInternalServerError).JSON(jsonResponse)
		}

		responseSpots := make([]dto.SpotListOut, 0)
		for _, s := range *spots {
			userSerializer := serializer.NewUserSerializer(&s.User)
			categorySerializer := serializer.NewCategorySerializer(&s.Category)
			spotSerializer := serializer.NewSpotSerializer(&s, userSerializer, categorySerializer)
			responseSpots = append(responseSpots, spotSerializer.ListSerialize(db))
		}

		jsonResponse := presenter.NewJsonResponse(false, "", responseSpots)
		return c.Status(http.StatusOK).JSON(jsonResponse)
	}
}
