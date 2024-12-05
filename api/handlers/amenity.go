package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/amenity"
	"github.com/gofiber/fiber/v2"
)

func CreateAmenity(service amenity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var jsonResponse presenter.JsonResponse
		var requestBody entities.CreateAmenityInput
		var serializer entities.AmenitySerializer

		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		createdAmenity, err := service.AddAmenity(&requestBody, c)
		if err != nil {
			jsonResponse = presenter.JsonResponse{
				Error:   true,
				Message: err.Error(),
				Data:    nil,
			}
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		serializer = createdAmenity
		jsonResponse = presenter.JsonResponse{
			Error:   false,
			Message: "",
			Data:    serializer.ListSerialize(),
		}

		return c.Status(fiber.StatusOK).JSON(jsonResponse)

	}
}
