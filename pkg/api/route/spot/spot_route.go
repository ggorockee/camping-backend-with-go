package spotroute

import (
	spotservice "camping-backend-with-go/internal/domain/service/spot"
	"camping-backend-with-go/pkg/api/middleware"

	spothandler "camping-backend-with-go/pkg/api/handler/spot"
	"github.com/gofiber/fiber/v2"
)

func SpotRouter(app fiber.Router, service spotservice.SpotService) {
	//public router
	publicSpotRouter := app.Group("/spot")
	publicSpotRouter.Get("/", spothandler.GetAllSpots(service))
	publicSpotRouter.Get("/:id/reviews", spothandler.SpotReviews(service))

	//private router
	// 로그인 인증이 되어야함
	privateSpotRouter := app.Group("/spot", middleware.Protected())
	privateSpotRouter.Post("/:id/review", spothandler.AddSpotReview(service))
	privateSpotRouter.Get("/:id", spothandler.GetSpot(service))
	privateSpotRouter.Put("/:id", spothandler.UpdateSpot(service))
	privateSpotRouter.Delete("/:id", spothandler.RemoveSpot(service))
	privateSpotRouter.Post("/", spothandler.AddSpot(service))
}
