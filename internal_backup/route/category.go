package route

import (
	"camping-backend-with-go/internal_backup/handler"
	"camping-backend-with-go/pkg/service/category"

	"github.com/gofiber/fiber/v2"
)

func CategoryRouter(app fiber.Router, service category.Service) {
	categoryRoute := app.Group("/category")
	//categoryRoute := app.Group("/category", middleware.Protected())

	categoryRoute.Get("/", handler.GetCategoryList(service))
	categoryRoute.Post("/", handler.CreateCategory(service))
	categoryRoute.Get("/:id", handler.GetCategory(service))
	categoryRoute.Put("/:id", handler.UpdateCategory(service))
	categoryRoute.Delete("/:id", handler.DeleteCategory(service))
}
