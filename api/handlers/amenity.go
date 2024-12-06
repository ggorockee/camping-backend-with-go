package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/amenity"
	"github.com/gofiber/fiber/v2"
)

// CreateAmenity
// @Summary CreateAmenity
// @Description CreateAmenity
// @Tags Amenity
// @Accept json
// @Produce json
// @Param amenity body entities.CreateAmenityInput true "Create Amenity"
// @Success 200 {object} presenter.JsonResponse{data=entities.AmenityDetailOut}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /amenity [post]
// @Security Bearer
func CreateAmenity(service amenity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.CreateAmenityInput

		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		createdAmenity, err := service.AddAmenity(&requestBody, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		//serializer = createdAmenity
		serializer := entities.NewAmenitySerializer(createdAmenity)
		jsonResponse := presenter.NewJsonResponse(false, "", serializer.ListSerialize())

		return c.Status(fiber.StatusOK).JSON(jsonResponse)

	}
}

// GetAmenities
// @Summary GetAmenities
// @Description GetAmenities
// @Tags Amenity
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{data=[]entities.AmenityListOut}
// @Failure 503 {object} presenter.JsonResponse{data=[]entities.AmenityListOut}
// @Router /amenity [get]
// @Security Bearer
func GetAmenities(service amenity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetchedAmenities, err := service.GetAmenities(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		serializedAmenities := make([]entities.AmenityListOut, 0)

		for _, fetchedAmenity := range *fetchedAmenities {
			serializer := entities.NewAmenitySerializer(&fetchedAmenity)
			serializedAmenities = append(serializedAmenities, serializer.ListSerialize())
		}

		jsonResponse := presenter.NewJsonResponse(false, "", serializedAmenities)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// GetAmenity
// @Summary GetAmenity
// @Description GetAmenity
// @Tags Amenity
// @Accept json
// @Produce json
// @Param id path int true "Amenity ID"
// @Success 200 {object} presenter.JsonResponse{data=entities.AmenityDetailOut}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /amenity/{id} [get]
// @Security Bearer
func GetAmenity(service amenity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		fetchedAmenity, err := service.GetAmenity(id, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}
		serializer := entities.NewAmenitySerializer(fetchedAmenity)
		jsonResponse := presenter.JsonResponse{
			Error:   false,
			Message: "",
			Data:    serializer.DetailSerialize(),
		}
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}
