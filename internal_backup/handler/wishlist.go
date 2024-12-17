package handler

import (
	"camping-backend-with-go/internal_backup/domain"
	"camping-backend-with-go/internal_backup/presenter"
	"camping-backend-with-go/internal_backup/service"
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/dto/wishdto"
	"camping-backend-with-go/pkg/serializer"
	"camping-backend-with-go/pkg/serializer/wishmarshal"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

// CreateWishList
// @Summary CreateWishList
// @Description CreateWishList
// @Tags WishList
// @Accept json
// @Produce json
// @Param wishlist body wishdto.CreateWishListReq true "Create WishList"
// @Success 200 {object} presenter.JsonResponse{data=wishmarshal.WishListRes}
// @Failure 503 {object} presenter.JsonResponse
// @Router /wishlist [post]
// @Security Bearer
func CreateWishList(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody wishdto.CreateWishListReq
		db, ok := c.Locals("db").(*gorm.DB)
		if !ok {
			log.Fatal("database load failed...")
		}
		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		wishList, err := service.CreateWishList(&requestBody, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		spotSerializer := func(spots []domain.Spot) []dto.SpotListOut {
			return serializer.SpotsSerializer(spots, db, c)
		}

		serializedWishList := wishmarshal.NewWishListSerializer(wishList).
			WishSpotsSerializer(spotSerializer)
		jsonResponse := presenter.NewJsonResponse(false, "", serializedWishList.Serialize())
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// GetWishList
// @Summary GetWishList
// @Description GetWishList
// @Tags WishList
// @Accept json
// @Produce json
// @Param id path int true "wishlist id"
// @Success 200 {object} presenter.JsonResponse{data=wishmarshal.WishListRes}
// @Failure 503 {object} presenter.JsonResponse
// @Router /wishlist/{id} [get]
// @Security Bearer
func GetWishList(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, ok := c.Locals("db").(*gorm.DB)
		if !ok {
			log.Fatal("database load failed...")
		}

		wishListId, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		wishList, err := service.GetWishList(wishListId, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		spotSerializer := func(spots []domain.Spot) []dto.SpotListOut {
			return serializer.SpotsSerializer(spots, db, c)
		}

		serializedWishList := wishmarshal.NewWishListSerializer(wishList).
			WishSpotsSerializer(spotSerializer)
		jsonResponse := presenter.NewJsonResponse(false, "", serializedWishList.Serialize())
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// UpdateWishList
// @Summary UpdateWishList
// @Description UpdateWishList
// @Tags WishList
// @Accept json
// @Produce json
// @Param id path int true "wishlist id"
// @Param wishlist body wishdto.UpdateWishListReq true "Update WishList"
// @Success 200 {object} presenter.JsonResponse{data=wishmarshal.WishListRes}
// @Failure 503 {object} presenter.JsonResponse
// @Router /wishlist/{id} [put]
// @Security Bearer
func UpdateWishList(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, ok := c.Locals("db").(*gorm.DB)
		if !ok {
			log.Fatal("database load failed...")
		}

		wishListId, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		var requestBody wishdto.UpdateWishListReq
		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		wishList, err := service.UpdateWishList(&requestBody, wishListId, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		spotSerializer := func(spots []domain.Spot) []dto.SpotListOut {
			return serializer.SpotsSerializer(spots, db, c)
		}

		serializedWishList := wishmarshal.NewWishListSerializer(wishList).
			WishSpotsSerializer(spotSerializer)
		jsonResponse := presenter.NewJsonResponse(false, "", serializedWishList.Serialize())
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// DeleteWishList
// @Summary DeleteWishList
// @Description DeleteWishList
// @Tags WishList
// @Accept json
// @Produce json
// @Param id path int true "wishlist id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /wishlist/{id} [delete]
// @Security Bearer
func DeleteWishList(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		wishListId, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		err = service.DeleteWishList(wishListId, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "successfull delete ", nil)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// GetWishLists
// @Summary GetWishLists
// @Description GetWishLists
// @Tags WishList
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{data=[]wishmarshal.WishListRes}
// @Failure 503 {object} presenter.JsonResponse
// @Router /wishlist [get]
// @Security Bearer
func GetWishLists(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		wishLists, err := service.GetWishLists(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		// user
		// category
		// amenitty
		// spot

		db := c.Locals("db").(*gorm.DB)
		spotsSerializer := func(spots []domain.Spot) []dto.SpotListOut {
			return serializer.SpotsSerializer(spots, db, c)
		}

		//userSerializer := func(user entities.User) dto.TinyUserOut {
		//	return helpers.SerializeUser(user)
		//}

		serializedWishLists := wishmarshal.NewWishListsSerializer(wishLists).
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
// @Success 200 {object} presenter.JsonResponse{data=[]wishmarshal.WishListRes}
// @Failure 503 {object} presenter.JsonResponse
// @Router /wishlist/{id}/spot/{spot_id} [put]
// @Security Bearer
func WishListToggle(service service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		wishListId, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		spotId, err := c.ParamsInt("spotId")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		err = service.WishListToggle(wishListId, spotId, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "added spot! to your wishlist", nil)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}
