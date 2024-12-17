package route

import (
	"camping-backend-with-go/internal_backup/handler"
	"camping-backend-with-go/internal_backup/middleware"
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/service/amenity"
	"github.com/gofiber/fiber/v2"
)

// api list
// /spot/amenities
// /spot/amenities/1

func AmenityRouter(app fiber.Router, service amenity.Service) {
	privateAmenityRoute := app.Group("/spot/amenity")
	//privateAmenityRoute := app.Group("/spot/amenity", middleware.Protected())
	privateAmenityRoute.Get("/", middleware.RoleMiddleware(
		dto.Staff,
		dto.Client,
		dto.Admin,
		dto.Owner,
	), handler.GetAmenities(service))
	privateAmenityRoute.Post("/", middleware.RoleMiddleware(
		dto.Staff,
		dto.Client,
		dto.Admin,
		dto.Owner,
	), handler.CreateAmenity(service))
	privateAmenityRoute.Get("/:id", middleware.RoleMiddleware(
		dto.Staff,
		dto.Client,
		dto.Admin,
		dto.Owner,
	), handler.GetAmenity(service))
	privateAmenityRoute.Put("/:id", middleware.RoleMiddleware(
		dto.Staff,
		dto.Client,
		dto.Admin,
		dto.Owner,
	), handler.UpdateAmenity(service))
	privateAmenityRoute.Delete("/:id", middleware.RoleMiddleware(
		dto.Staff,
		dto.Client,
		dto.Admin,
		dto.Owner,
	), handler.DeleteAmenity(service))

}
