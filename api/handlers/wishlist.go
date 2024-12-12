package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/serializer"
	"camping-backend-with-go/pkg/serializer/helpers"
	"camping-backend-with-go/pkg/service/wishlistsvc"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetWishList
// @Summary GetWishList
// @Description GetWishList
// @Tags WishList
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{data=[]serializer.WishListRes}
// @Failure 503 {object} presenter.JsonResponse
// @Router /wishlist [get]
// @Security Bearer
func GetWishList(controller wishlistsvc.Controller) fiber.Handler {
	return func(c *fiber.Ctx) error {
		wishList, err := controller.GetWishList(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		// user
		// category
		// amenitty
		// spot

		db := c.Locals("db").(*gorm.DB)
		spotSerializer := func(spots []entities.Spot) []dto.SpotListOut {
			return helpers.SerializeSpot(spots, db, c)
		}

		//userSerializer := func(user entities.User) dto.TinyUserOut {
		//	return helpers.SerializeUser(user)
		//}

		serializedWishList := serializer.NewWishListSerializer(wishList).
			WithSpotSerializer(spotSerializer)

		jsonResponse := presenter.NewJsonResponse(false, "", serializedWishList.Serialize())
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}
