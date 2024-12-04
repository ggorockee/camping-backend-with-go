package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/api/middleware"
	"camping-backend-with-go/pkg/service/spot"

	"github.com/gofiber/fiber/v2"
)

func SpotRouter(app fiber.Router, service spot.Service) {
	app.Get("/spot/me", middleware.Protected(), handlers.GetMySpots(service))
	app.Get("/spot/:id", middleware.Protected(), handlers.GetSpot(service))
	app.Put("/spot/:id", middleware.Protected(), handlers.UpdateSpot(service))
	app.Patch("/spot/:id", handlers.PartialUpdateSpot(service))
	app.Post("/spot", middleware.Protected(), handlers.AddSpot(service))
	app.Delete("/spot/:id", handlers.RemoveSpot(service))
}
