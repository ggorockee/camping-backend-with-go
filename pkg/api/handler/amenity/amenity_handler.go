package amenityhandler

import (
	amenitydto "camping-backend-with-go/internal/application/dto/amenity"
	"camping-backend-with-go/internal/domain/presenter"
	amenityservice "camping-backend-with-go/internal/domain/service/amenity"
	"github.com/gofiber/fiber/v2"
	"log"
)

// CreateAmenity
// @Summary CreateAmenity
// @Description CreateAmenity
// @Tags Amenity
// @Accept json
// @Produce json
// @Param requestBody body amenitydto.CreateAmenityReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/amenity [post]
// @Security Bearer
func CreateAmenity(service amenityservice.AmenityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody amenitydto.CreateAmenityReq

		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		amenity, err := service.CreateAmenity(&requestBody, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		//serializer = createdAmenity
		//AmenitySerializer := serializer.NewAmenitySerializer(createdAmenity)
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Println(">>>> serialize not implemented")
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		jsonResponse := presenter.NewJsonResponse(false, "", amenity)

		return c.Status(fiber.StatusOK).JSON(jsonResponse)

	}
}

// GetAmenities
// @Summary GetAmenities
// @Description GetAmenities
// @Tags Amenity
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/amenity [get]
// @Security Bearer
func GetAmenities(service amenityservice.AmenityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		amenities, err := service.GetAmenityList(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		//serializedAmenities := make([]dto.AmenityListOut, 0)
		//
		//for _, fetchedAmenity := range *fetchedAmenities {
		//	AmenitySerializer := serializer.NewAmenitySerializer(&fetchedAmenity)
		//	serializedAmenities = append(serializedAmenities, AmenitySerializer.ListSerialize())
		//}
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Println(">>>> serialize not implemented")
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		jsonResponse := presenter.NewJsonResponse(false, "", amenities)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// GetAmenity
// @Summary GetAmenity
// @Description GetAmenity
// @Tags Amenity
// @Accept json
// @Produce json
// @Param id path int true "amenity id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/amenity/{id} [get]
// @Security Bearer
func GetAmenity(service amenityservice.AmenityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		amenity, err := service.GetAmenityById(id)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}
		//AmenitySerializer := serializer.NewAmenitySerializer(fetchedAmenity)
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Println(">>>> serialize not implemented")
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		jsonResponse := presenter.NewJsonResponse(false, "", amenity)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// UpdateAmenity
// @Summary UpdateAmenity
// @Description UpdateAmenity
// @Tags Amenity
// @Accept json
// @Produce json
// @Param id path int true "amenity id"
// @Param requestBody body amenitydto.UpdateAmenityReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/amenity/{id} [put]
// @Security Bearer
func UpdateAmenity(service amenityservice.AmenityService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody amenitydto.UpdateAmenityReq
		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), "")
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		id, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		amenity, err := service.UpdateAmenity(&requestBody, id, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		//AmenitySerializer := serializer.NewAmenitySerializer(updatedAmenity)
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Println(">>>> serialize not implemented")
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		jsonResponse := presenter.NewJsonResponse(false, "", amenity)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// DeleteAmenity
// @Summary DeleteAmenity
// @Description DeleteAmenity
// @Tags Amenity
// @Accept json
// @Produce json
// @Param id path int true "amenity id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/amenity/{id} [delete]
// @Security Bearer
func DeleteAmenity(service amenityservice.AmenityService) fiber.Handler {
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
