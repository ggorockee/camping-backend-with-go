package route

import (
	"camping-backend-with-go/internal/handler"
	"camping-backend-with-go/internal/middleware"
	"camping-backend-with-go/internal/service"
	"github.com/gofiber/fiber/v2"
)

func WishListRouter(app fiber.Router, service service.Service) {
	wishlistRouter := app.Group("/wishlist", middleware.Protected())
	wishlistRouter.Get("/", handler.GetWishLists(service))
	wishlistRouter.Post("/", handler.CreateWishList(service))
	wishlistRouter.Get("/:id", handler.GetWishList(service))
	wishlistRouter.Put("/:id", handler.UpdateWishList(service))
	wishlistRouter.Put("/:id", handler.DeleteWishList(service))
	wishlistRouter.Put("/:id/spot/:spotId", handler.WishListToggle(service))
}
