package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/pkg/spot"
	"github.com/gofiber/fiber/v2"
)

func SpotRouter(app fiber.Router, service spot.Service) {
	app.Get("/spots", handlers.GetSpots(service))
	app.Get("/spots/:id", handlers.GetSpot(service))
	app.Put("/spots/:id", handlers.UpdateSpot(service))
	app.Patch("/spots/:id", handlers.PartialUpdateSpot(service))
	app.Post("/spots", handlers.AddSpot(service))
	app.Delete("/spots/:id", handlers.RemoveSpot(service))
}
