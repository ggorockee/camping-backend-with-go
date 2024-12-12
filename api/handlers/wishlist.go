package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/serializer"
	"camping-backend-with-go/pkg/service/wishlistsvc"

	"github.com/gofiber/fiber/v2"
)

func GetWishList(controller wishlistsvc.Controller) fiber.Handler {
	return func(c *fiber.Ctx) error {
		wishList, err := controller.GetWishList(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		serializedWishList := serializer.NewWishListSerializer(wishList)
		jsonResponse := presenter.NewJsonResponse(false, "", serializedWishList.Serialize())
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}
