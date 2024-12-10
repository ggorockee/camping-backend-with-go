package routes

import (
	"camping-backend-with-go/api/handlers"
	"camping-backend-with-go/pkg/service/category"

	"github.com/gofiber/fiber/v2"
)

func CategoryRouter(app fiber.Router, service category.Service) {
	categoryRoute := app.Group("/category")
	//categoryRoute := app.Group("/category", middleware.Protected())

	categoryRoute.Get("/", handlers.GetCategoryList(service))
	categoryRoute.Post("/", handlers.CreateCategory(service))
	categoryRoute.Get("/:id", handlers.GetCategory(service))
	categoryRoute.Put("/:id", handlers.UpdateCategory(service))
	categoryRoute.Delete("/:id", handlers.DeleteCategory(service))
}
