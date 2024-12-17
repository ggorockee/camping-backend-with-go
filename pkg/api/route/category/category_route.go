package categoryroute

import (
	categoryservice "camping-backend-with-go/internal/domain/service/category"
	categoryhandler "camping-backend-with-go/pkg/api/handler/category"
	"github.com/gofiber/fiber/v2"
)

func CategoryRouter(app fiber.Router, service categoryservice.CategoryService) {
	categoryRouter := app.Group("/category")
	categoryRouter.Get("/", categoryhandler.GetCategoryList(service))
	categoryRouter.Post("/", categoryhandler.CreateCategory(service))
	categoryRouter.Get("/:id", categoryhandler.GetCategory(service))
	categoryRouter.Put("/:id", categoryhandler.UpdateCategory(service))
	categoryRouter.Delete("/:id", categoryhandler.DeleteCategory(service))
}
