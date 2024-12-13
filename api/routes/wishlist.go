package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/api/middleware"
	"camping-backend-with-go/pkg/service/wishlist"

	"github.com/gofiber/fiber/v2"
)

func WishListRouter(app fiber.Router, service wishlist.Service) {
	wishlistRouter := app.Group("/wishlist", middleware.Protected())
	wishlistRouter.Get("/", handlers.GetWishLists(service))
	wishlistRouter.Post("/", handlers.CreateWishList(service))
	wishlistRouter.Get("/:id", handlers.GetWishList(service))
	wishlistRouter.Put("/:id", handlers.UpdateWishList(service))
	wishlistRouter.Put("/:id", handlers.DeleteWishList(service))
	wishlistRouter.Put("/:id/spot/:spotId", handlers.WishListToggle(service))
}
