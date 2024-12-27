package categoryhandler

import (
	"camping-backend-with-go/internal/application/dto"

	"camping-backend-with-go/internal/domain/presenter"
	categoryservice "camping-backend-with-go/internal/domain/service/category"

	"github.com/gofiber/fiber/v2"
)

// GetCategoryList
// @Summary GetCategoryList
// @Description GetCategoryList
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /category [get]
// @Security Bearer
func GetCategoryList(service categoryservice.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		categories, err := service.GetCategoryList(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "", categories)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// CreateCategory
// @Summary CreateCategory
// @Description CreateCategory
// @Tags Category
// @Accept json
// @Produce json
// @Param requestBody body dto.CreateCategoryReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /category [post]
// @Security Bearer
func CreateCategory(service categoryservice.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.CreateCategoryReq

		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		category, err := service.CreateCategory(&requestBody, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		// ser := serializer.New(category, commonhandler.SerializerFactory)
		jsonResponse := presenter.NewJsonResponse(false, "", category)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// GetCategory
// @Summary GetCategory
// @Description GetCategory
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "category id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /category/{id} [get]
// @Security Bearer
func GetCategory(service categoryservice.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		categoryId := c.Params("id", "")
		if categoryId == "" {
			jsonResponse := presenter.NewJsonResponse(true, "id parsing 실패", nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		category, err := service.GetCategoryById(categoryId, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		// ser := serializer.New(category, commonhandler.SerializerFactory)
		jsonResponse := presenter.NewJsonResponse(false, "", category)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// UpdateCategory
// @Summary UpdateCategory
// @Description UpdateCategory
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "category id"
// @Param requestBody body dto.UpdateCategoryReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /category/{id} [put]
// @Security Bearer
func UpdateCategory(service categoryservice.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.UpdateCategoryReq

		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		categoryId := c.Params("id", "")
		if categoryId == "" {
			jsonResponse := presenter.NewJsonResponse(true, "id parsing 실패", nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		category, err := service.UpdateCategory(&requestBody, categoryId, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		// ser := serializer.New(category, commonhandler.SerializerFactory)
		jsonResponse := presenter.NewJsonResponse(false, "", category)

		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// DeleteCategory
// @Summary DeleteCategory
// @Description DeleteCategory
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "category id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /category/{id} [delete]
// @Security Bearer
func DeleteCategory(service categoryservice.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		categoryId := c.Params("id", "")
		if categoryId == "" {
			jsonResponse := presenter.NewJsonResponse(true, "id parsing 실패", nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "Successful delete", nil)

		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}
