package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/pkg/spot"
	"github.com/gofiber/fiber/v2"
)

func SpotRouter(app fiber.Router, service spot.Service) {
	app.Get("/spot", handlers.GetSpots(service))
	app.Get("/spot/:id", handlers.GetSpot(service))
	app.Put("/spot/:id", handlers.UpdateSpot(service))
	app.Patch("/spot/:id", handlers.PartialUpdateSpot(service))
	app.Post("/spot", handlers.AddSpot(service))
	app.Delete("/spot/:id", handlers.RemoveSpot(service))
}
