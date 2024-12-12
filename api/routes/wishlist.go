package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/api/middleware"
	"camping-backend-with-go/pkg/service/wishlistsvc"

	"github.com/gofiber/fiber/v2"
)

func WishListRouter(app fiber.Router, controller wishlistsvc.Controller) {
	wishlistRouter := app.Group("/wishlist", middleware.Protected())
	wishlistRouter.Get("/", handlers.GetWishList(controller))
	wishlistRouter.Put("/:id/spot/:spotId", handlers.WishListToggle(controller))
}
