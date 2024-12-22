package spothandler

import (
	reviewdto "camping-backend-with-go/internal/application/dto/review"
	spotdto "camping-backend-with-go/internal/application/dto/spot"

	"camping-backend-with-go/internal/domain/presenter"
	spotservice "camping-backend-with-go/internal/domain/service/spot"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// AddSpotReview
// @Summary AddSpotReview
// @Description AddSpotReview
// @Tags Spot
// @Accept json
// @Produce json
// @Param requestBody body reviewdto.CreateSpotReviewReq true "requestBody"
// @Param id path string true "spot id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/{id}/review [post]
// @Security Bearer
func AddSpotReview(service spotservice.SpotService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody reviewdto.CreateSpotReviewReq
		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		spotId := c.Params("id", "")
		if spotId == "" {
			jsonResponse := presenter.NewJsonResponse(true, "id parsing 실패", nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		spot, err := service.GetSpotById(spotId, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		review, err := service.CreateSpotReview(&requestBody, spot, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)

		}

		// ser := serializer.New(review, commonhandler.SerializerFactory)

		jsonResponse := presenter.NewJsonResponse(false, "", review)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// AddSpot
// @Summary AddSpot
// @Description AddSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param requestBody body spotdto.CreateSpotReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot [post]
// @Security Bearer
func AddSpot(service spotservice.SpotService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody spotdto.CreateSpotReq
		db := c.Locals("db").(*gorm.DB)
		log.Printf("db: %v\n", db)

		err := c.BodyParser(&requestBody)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		spot, err := service.CreateSpot(&requestBody, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}
		log.Println(">>> handler", spot.User)

		jsonResponse := presenter.NewJsonResponse(true, "", spot)

		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// UpdateSpot
// @Summary UpdateSpot
// @Description UpdateSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param id path string true "spot id"
// @Param requestBody body spotdto.UpdateSpotReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/{id} [put]
// @Security Bearer
func UpdateSpot(service spotservice.SpotService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		log.Printf("db %v\n", db)
		var requestBody spotdto.UpdateSpotReq

		err := c.BodyParser(&requestBody)
		spotId := c.Params("id", "")
		if spotId == "" {
			jsonResponse := presenter.NewJsonResponse(true, "id parsing 실패", nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		spot, err := service.UpdateSpot(&requestBody, spotId, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Println(">>>> serialize not implemented")
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		//userSerializer := serializer.NewUserSerializer(&fetchedSpot.User)
		//categorySerializer := serializer.NewCategorySerializer(&fetchedSpot.Category)
		//spotSerializer := serializer.NewSpotSerializer(fetchedSpot, userSerializer, categorySerializer)
		// var serializedSpot response.SpotDetailRes
		// err = serializer.GeneralSerializer(spot, &serializedSpot)

		jsonResponse := presenter.NewJsonResponse(false, "", spot)

		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// GetSpot
// @Summary GetSpot
// @Description GetSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param id path string true "spot id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/{id} [get]
// @Security Bearer
func GetSpot(service spotservice.SpotService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the spot ID from the request parameters
		spotId := c.Params("id", "")
		if spotId == "" {
			// Return an error response if the ID is missing
			jsonResponse := presenter.NewJsonResponse(true, "Missing spot ID", nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		// Fetch the spot from the service
		spot, err := service.GetSpotById(spotId, c)
		if err != nil {
			// Return an error response if the spot couldn't be fetched
			jsonResponse := presenter.NewJsonResponse(true, "Failed to fetch spot: "+err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		// Create a new SpotDetailSerializer
		// spotSerializer := &spotserializer.SpotDetailSerializer{Spot: spot}

		// Serialize the spot
		// serializedSpot := spotSerializer.Serialize()

		// Create and return the JSON response
		jsonResponse := presenter.NewJsonResponse(false, "", spot)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// RemoveSpot
// @Summary RemoveSpot
// @Description RemoveSpot
// @Tags Spot
// @Accept json
// @Produce json
// @Param id path string true "spot id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/{id} [delete]
// @Security Bearer
func RemoveSpot(service spotservice.SpotService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		spotId := c.Params("id", "")
		if spotId == "" {
			jsonResponse := presenter.NewJsonResponse(true, "id parsing 실패", nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}
		err := service.DeleteSpot(spotId, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "successfully delete", nil)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// SpotReviews
// @Summary SpotReviews
// @Description SpotReviews
// @Tags Spot
// @Accept json
// @Produce json
// @Param id path string true "spot id"
// @Param page query int false "Page number" default(1)
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot/{id}/reviews [get]
// @Security Bearer
func SpotReviews(service spotservice.SpotService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		spotId := c.Params("id", "")
		if spotId == "" {
			jsonResponse := presenter.NewJsonResponse(true, "id parsing 실패", nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		spot, err := service.GetSpotById(spotId, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		reviews, err := service.GetReviewsFromSpot(spot, c)

		// query params
		pageStr := c.Query("page", "1")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
		}

		pageSize := 5
		pageStart := (page - 1) * pageSize
		pageEnd := pageStart + pageSize
		log.Println("pageEnd: ", pageEnd)

		//reviewsSerializer := serializer.NewReviewsSerializer(reviews)
		//jsonResponse := presenter.NewJsonResponse(false, "", reviewsSerializer.Serializer()[pageStart:pageEnd])
		//serializedReviews := reviewsSerializer.Serialize()
		//paginatedReviews := make([]serializer.ReviewOut, 0)
		//switch {
		//case len(serializedReviews) == 0:
		//case len(serializedReviews) <= pageStart:
		//	// pageStart가 전체 리뷰 수를 초과할 경우 마지막 페이지 반환
		//	lastPageStart := (len(serializedReviews) - 1) / pageSize * pageSize
		//	paginatedReviews = serializedReviews[lastPageStart:]
		//case len(serializedReviews) < pageEnd:
		//	paginatedReviews = serializedReviews[pageStart:]
		//default:
		//	paginatedReviews = serializedReviews[pageStart:pageEnd]
		//}

		jsonResponse := presenter.NewJsonResponse(false, "", reviews)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// GetAllSpots
// @Summary GetAllSpots
// @Description GetAllSpots
// @Tags Spot
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /spot [get]
// @Security Bearer
func GetAllSpots(service spotservice.SpotService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := c.Locals("db").(*gorm.DB)
		log.Printf("db: %v\n", db)
		spots, err := service.GetAllSpots() //*[]entities.Spot

		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		//responseSpots := make([]dto.SpotListOut, 0)
		//for _, s := range *spots {
		//	userSerializer := serializer.NewUserSerializer(&s.User)
		//	categorySerializer := serializer.NewCategorySerializer(&s.Category)
		//	spotSerializer := serializer.NewSpotSerializer(&s, userSerializer, categorySerializer)
		//	responseSpots = append(responseSpots, spotSerializer.ListSerialize(db, c))
		//}

		jsonResponse := presenter.NewJsonResponse(false, "", spots)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}
