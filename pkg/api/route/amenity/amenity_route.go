package amenityroute

import (
	amenityservice "camping-backend-with-go/internal/domain/service/amenity"
	amenityhandler "camping-backend-with-go/pkg/api/handler/amenity"
	"github.com/gofiber/fiber/v2"
)

func AmenityRouter(app fiber.Router, service amenityservice.AmenityService) {
	amenityRouter := app.Group("/spot/amenity")
	amenityRouter.Get("/", amenityhandler.GetAmenities(service))
	amenityRouter.Post("/", amenityhandler.CreateAmenity(service))
	amenityRouter.Get("/:id", amenityhandler.GetAmenity(service))
	amenityRouter.Put("/:id", amenityhandler.UpdateAmenity(service))
	amenityRouter.Delete("/:id", amenityhandler.DeleteAmenity(service))
}
