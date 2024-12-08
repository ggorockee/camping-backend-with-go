package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/api/middleware"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/amenity"
	"github.com/gofiber/fiber/v2"
)

// api list
// /spot/amenities
// /spot/amenities/1

func AmenityRouter(app fiber.Router, service amenity.Service) {
	privateAmenityRoute := app.Group("/amenity", middleware.Protected())
	privateAmenityRoute.Get("/", middleware.RoleMiddleware(
		entities.Staff,
		entities.Client,
		entities.Admin,
		entities.Owner,
	), handlers.GetAmenities(service))
	privateAmenityRoute.Post("/", middleware.RoleMiddleware(
		entities.Staff,
		entities.Client,
		entities.Admin,
		entities.Owner,
	), handlers.CreateAmenity(service))
	privateAmenityRoute.Get("/:id", middleware.RoleMiddleware(
		entities.Staff,
		entities.Client,
		entities.Admin,
		entities.Owner,
	), handlers.GetAmenity(service))
	privateAmenityRoute.Put("/:id", middleware.RoleMiddleware(
		entities.Staff,
		entities.Client,
		entities.Admin,
		entities.Owner,
	), handlers.UpdateAmenity(service))
	privateAmenityRoute.Delete("/:id", middleware.RoleMiddleware(
		entities.Staff,
		entities.Client,
		entities.Admin,
		entities.Owner,
	), handlers.DeleteAmenity(service))

}
