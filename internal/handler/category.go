package handler

import (
	"camping-backend-with-go/internal/presenter"
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/serializer"
	"camping-backend-with-go/pkg/service/category"

	"github.com/gofiber/fiber/v2"
)

// GetCategoryList is a function to get category data from database
// @Summary GetCategoryList
// @Description GetCategoryList
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} presenter.JsonResponse{data=[]dto.CategoryListOut}
// @Failure 503 {object} presenter.JsonResponse
// @Router /category [get]
// @Security Bearer
func GetCategoryList(service category.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetchedCategories, err := service.GetCategoryList(c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		serializedCategories := make([]dto.CategoryListOut, 0)

		for _, fetchedCategory := range *fetchedCategories {
			categorySerializer := serializer.NewCategorySerializer(&fetchedCategory)
			serializedCategories = append(serializedCategories, categorySerializer.ListSerialize())
		}

		jsonResponse := presenter.NewJsonResponse(false, "", serializedCategories)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// CreateCategory is a function to create category data to database
// @Summary CreateCategory
// @Description CreateCategory
// @Tags Category
// @Accept json
// @Produce json
// @Param user body dto.CreateCategoryIn true "Create Category Schema"
// @Success 200 {object} presenter.JsonResponse{data=dto.CategoryListOut}
// @Failure 503 {object} presenter.JsonResponse
// @Router /category [post]
// @Security Bearer
func CreateCategory(service category.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.CreateCategoryIn

		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		createdCategory, err := service.CreateCategory(&requestBody, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		categorySerializer := serializer.NewCategorySerializer(createdCategory)
		jsonResponse := presenter.NewJsonResponse(false, "", categorySerializer.ListSerialize())

		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// GetCategory
// @Summary GetCategory
// @Description GetCategory
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category Id"
// @Success 200 {object} presenter.JsonResponse{data=dto.CategoryDetailOut}
// @Failure 503 {object} presenter.JsonResponse
// @Router /category/{id} [get]
// @Security Bearer
func GetCategory(service category.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		fetchedCategory, err := service.GetCategory(id, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		categorySerializer := serializer.NewCategorySerializer(fetchedCategory)
		jsonResponse := presenter.NewJsonResponse(false, "", categorySerializer.DetailSerialize())

		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// UpdateCategory
// @Summary UpdateCategory
// @Description UpdateCategory
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category Id"
// @Param user body dto.UpdateCategoryIn true "Update Category"
// @Success 200 {object} presenter.JsonResponse{data=dto.CategoryDetailOut}
// @Failure 503 {object} presenter.JsonResponse
// @Router /category/{id} [put]
// @Security Bearer
func UpdateCategory(service category.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.UpdateCategoryIn

		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		id, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		fetchedCategory, err := service.UpdateCategory(&requestBody, id, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		categorySerializer := serializer.NewCategorySerializer(fetchedCategory)
		jsonResponse := presenter.NewJsonResponse(false, "", categorySerializer.DetailSerialize())

		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// DeleteCategory
// @Summary DeleteCategory
// @Description DeleteCategory
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category Id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /category/{id} [delete]
// @Security Bearer
func DeleteCategory(service category.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		err = service.DeleteCategory(id, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		jsonResponse := presenter.NewJsonResponse(false, "Successful delete", nil)

		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}
