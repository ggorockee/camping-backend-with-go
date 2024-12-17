package route

import (
	"camping-backend-with-go/internal_backup/handler"
	"camping-backend-with-go/internal_backup/middleware"
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/service/spot"

	"github.com/gofiber/fiber/v2"
)

func SpotRouter(app fiber.Router, service spot.Service) {
	//public router
	publicSpotRouter := app.Group("/spot")
	publicSpotRouter.Get("/", handler.GetAllSpots(service))
	publicSpotRouter.Get("/:id/reviews", handler.SpotReviews(service))

	//private router
	// 로그인 인증이 되어야함
	privateSpotRouter := app.Group("/spot", middleware.Protected())
	//privateSpotRouter := app.Group("/spot", middleware.Protected())

	privateSpotRouter.Post("/:id/review", middleware.RoleMiddleware(
		dto.Client,
		dto.Owner,
		dto.Admin,
		dto.Staff,
	), handler.AddSpotReview(service))

	privateSpotRouter.Get("/:id", middleware.RoleMiddleware(
		dto.Client,
		dto.Owner,
		dto.Admin,
		dto.Staff,
	), handler.GetSpot(service))

	privateSpotRouter.Put("/:id", middleware.RoleMiddleware(
		dto.Client,
		dto.Owner,
		dto.Admin,
		dto.Staff,
	), handler.UpdateSpot(service))
	privateSpotRouter.Delete("/:id", middleware.RoleMiddleware(
		dto.Client,
		dto.Owner,
		dto.Admin,
		dto.Staff,
	), handler.RemoveSpot(service))
	privateSpotRouter.Post("/", middleware.RoleMiddleware(
		dto.Client,
		dto.Owner,
		dto.Admin,
		dto.Staff,
	), handler.AddSpot(service))
}
