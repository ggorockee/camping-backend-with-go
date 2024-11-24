package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/pkg/spot"
	"github.com/gofiber/fiber/v2"
)

func SpotRouter(app fiber.Router, service spot.Service) {
	app.Get("/spots", handlers.GetSpots(service))
	app.Post("/spots", handlers.AddSpot(service))
	//app.Put("/spots", handlers.UpdateSpot(service))
	//app.Delete("/spots", handlers.RemoveSpot(service))
}
