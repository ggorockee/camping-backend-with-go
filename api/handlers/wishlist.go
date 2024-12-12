package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/serializer"
	"camping-backend-with-go/pkg/service/wishlistsvc"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetWishLists
// @Summary GetWishLists
// @Description GetWishLists
// @Tags WishList
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{data=[]serializer.WishListRes}
// @Failure 503 {object} presenter.JsonResponse
// @Router /wishlists [get]
// @Security Bearer
func GetWishLists(controller wishlistsvc.Controller) fiber.Handler {
	return func(c *fiber.Ctx) error {
		wishLists, err := controller.GetWishLists(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		// user
		// category
		// amenitty
		// spot

		db := c.Locals("db").(*gorm.DB)
		spotsSerializer := func(spots []entities.Spot) []dto.SpotListOut {
			return serializer.SpotsSerializer(spots, db, c)
		}

		//userSerializer := func(user entities.User) dto.TinyUserOut {
		//	return helpers.SerializeUser(user)
		//}

		serializedWishLists := serializer.NewWishListsSerializer(wishLists).
			WithSpotsSerializer(spotsSerializer)

		jsonResponse := presenter.NewJsonResponse(false, "", serializedWishLists.Serialize())
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// WishListToggle
// @Summary WishListToggle
// @Description WishListToggle
// @Tags WishList
// @Accept json
// @Produce json
// @Param id path int true "wishlist id"
// @Param spot_id path int true "spot id"
// @Success 200 {object} presenter.JsonResponse{data=[]serializer.WishListRes}
// @Failure 503 {object} presenter.JsonResponse
// @Router /wishlist/{id}/spot/{spot_id} [put]
// @Security Bearer
func WishListToggle(controller wishlistsvc.Controller) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, ok := c.Locals("db").(*gorm.DB)
		if !ok {
			jsonResponse := presenter.NewJsonResponse(true, "failed Load database", nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}
		wishList, err := controller.WishListToggle(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		spotSerializer := func(spots []entities.Spot) []dto.SpotListOut {
			return serializer.SpotsSerializer(spots, db, c)
		}

		serializedWishList := serializer.NewWishListSerializer(wishList).
			WishSpotsSerializer(spotSerializer)

		jsonResponse := presenter.NewJsonResponse(false, "", serializedWishList.Serialize())
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}
