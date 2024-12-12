package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/api/middleware"
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/service/spot"

	"github.com/gofiber/fiber/v2"
)

func SpotRouter(app fiber.Router, service spot.Service) {
	//public router
	publicSpotRouter := app.Group("/spot")
	publicSpotRouter.Get("/", handlers.GetAllSpots(service))
	publicSpotRouter.Get("/:id/reviews", handlers.SpotReviews(service))

	//private router
	// 로그인 인증이 되어야함
	privateSpotRouter := app.Group("/spot", middleware.Protected())
	//privateSpotRouter := app.Group("/spot", middleware.Protected())

	privateSpotRouter.Post("/:id/review", middleware.RoleMiddleware(
		dto.Client,
		dto.Owner,
		dto.Admin,
		dto.Staff,
	), handlers.AddSpotReview(service))

	privateSpotRouter.Get("/:id", middleware.RoleMiddleware(
		dto.Client,
		dto.Owner,
		dto.Admin,
		dto.Staff,
	), handlers.GetSpot(service))

	privateSpotRouter.Put("/:id", middleware.RoleMiddleware(
		dto.Client,
		dto.Owner,
		dto.Admin,
		dto.Staff,
	), handlers.UpdateSpot(service))
	privateSpotRouter.Delete("/:id", middleware.RoleMiddleware(
		dto.Client,
		dto.Owner,
		dto.Admin,
		dto.Staff,
	), handlers.RemoveSpot(service))
	privateSpotRouter.Post("/", middleware.RoleMiddleware(
		dto.Client,
		dto.Owner,
		dto.Admin,
		dto.Staff,
	), handlers.AddSpot(service))
}
