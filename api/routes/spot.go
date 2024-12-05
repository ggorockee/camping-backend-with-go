package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/api/middleware"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/spot"

	"github.com/gofiber/fiber/v2"
)

func SpotRouter(app fiber.Router, service spot.Service) {
	//public router
	publicSpotRouter := app.Group("/spot")
	publicSpotRouter.Get("/spot", handlers.GetAllSpots(service))

	//private router
	privateSpotRouter := app.Group("/spot", middleware.Protected())
	privateSpotRouter.Get("/me", middleware.RoleMiddleware(
		entities.Client,
		entities.Owner,
		entities.Admin,
		entities.Staff,
	), handlers.GetMySpots(service))
	privateSpotRouter.Get("/:id", middleware.RoleMiddleware(
		entities.Client,
		entities.Owner,
		entities.Admin,
		entities.Staff,
	), handlers.GetSpot(service))
	privateSpotRouter.Put("/:id", middleware.RoleMiddleware(
		entities.Client,
		entities.Owner,
		entities.Admin,
		entities.Staff,
	), handlers.UpdateSpot(service))
	privateSpotRouter.Delete("/:id", middleware.RoleMiddleware(
		entities.Client,
		entities.Owner,
		entities.Admin,
		entities.Staff,
	), handlers.RemoveSpot(service))
	privateSpotRouter.Post("/", middleware.RoleMiddleware(
		entities.Client,
		entities.Owner,
		entities.Admin,
		entities.Staff,
	), handlers.AddSpot(service))
}
