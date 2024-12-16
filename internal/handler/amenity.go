package handler

import (
	"camping-backend-with-go/internal/presenter"
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/serializer"
	"camping-backend-with-go/pkg/service/amenity"
	"github.com/gofiber/fiber/v2"
)

// CreateAmenity
// @Summary CreateAmenity
// @Description CreateAmenity
// @Tags Amenity
// @Accept json
// @Produce json
// @Param amenity body dto.CreateAmenityIn true "Create Amenity"
// @Success 200 {object} presenter.JsonResponse{data=dto.AmenityDetailOut}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/amenity [post]
// @Security Bearer
func CreateAmenity(service amenity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.CreateAmenityIn

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
		AmenitySerializer := serializer.NewAmenitySerializer(createdAmenity)
		jsonResponse := presenter.NewJsonResponse(false, "", AmenitySerializer.ListSerialize())

		return c.Status(fiber.StatusOK).JSON(jsonResponse)

	}
}

// GetAmenities
// @Summary GetAmenities
// @Description GetAmenities
// @Tags Amenity
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{data=[]dto.AmenityListOut}
// @Failure 503 {object} presenter.JsonResponse{data=[]dto.AmenityListOut}
// @Router /spot/amenity [get]
// @Security Bearer
func GetAmenities(service amenity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetchedAmenities, err := service.GetAmenities(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		serializedAmenities := make([]dto.AmenityListOut, 0)

		for _, fetchedAmenity := range *fetchedAmenities {
			AmenitySerializer := serializer.NewAmenitySerializer(&fetchedAmenity)
			serializedAmenities = append(serializedAmenities, AmenitySerializer.ListSerialize())
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
// @Success 200 {object} presenter.JsonResponse{data=dto.AmenityDetailOut}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/amenity/{id} [get]
// @Security Bearer
func GetAmenity(service amenity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		fetchedAmenity, err := service.GetAmenity(id)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}
		AmenitySerializer := serializer.NewAmenitySerializer(fetchedAmenity)
		jsonResponse := presenter.JsonResponse{
			Error:   false,
			Message: "",
			Data:    AmenitySerializer.DetailSerialize(),
		}
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// UpdateAmenity
// @Summary UpdateAmenity
// @Description UpdateAmenity
// @Tags Amenity
// @Accept json
// @Produce json
// @Param id path int true "Amenity ID"
// @Param amenity body dto.UpdateAmenityIn true "Update Amenity"
// @Success 200 {object} presenter.JsonResponse{data=dto.AmenityDetailOut}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/amenity/{id} [put]
// @Security Bearer
func UpdateAmenity(service amenity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody dto.UpdateAmenityIn
		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), "")
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		id, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		updatedAmenity, err := service.UpdateAmenity(&requestBody, id, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		AmenitySerializer := serializer.NewAmenitySerializer(updatedAmenity)
		jsonResponse := presenter.NewJsonResponse(false, "", AmenitySerializer.DetailSerialize())
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// DeleteAmenity
// @Summary DeleteAmenity
// @Description DeleteAmenity
// @Tags Amenity
// @Accept json
// @Produce json
// @Param id path int true "Amenity ID"
// @Success 200 {object} presenter.JsonResponse{data=dto.AmenityDetailOut}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/amenity/{id} [delete]
// @Security Bearer
func DeleteAmenity(service amenity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		err = service.DeleteAmenity(id, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(true, "Deleted successfully", nil)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}
