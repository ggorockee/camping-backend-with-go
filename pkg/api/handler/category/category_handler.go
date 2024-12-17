package categoryhandler

import (
	categorydto "camping-backend-with-go/internal/application/dto/category"
	"camping-backend-with-go/internal/domain/presenter"
	categoryservice "camping-backend-with-go/internal/domain/service/category"
	//"camping-backend-with-go/pkg_backup/service/category"
	"github.com/gofiber/fiber/v2"
	"log"
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

		//serializedCategories := make([]dto.CategoryListOut, 0)
		//
		//for _, fetchedCategory := range *fetchedCategories {
		//	categorySerializer := serializer.NewCategorySerializer(&fetchedCategory)
		//	serializedCategories = append(serializedCategories, categorySerializer.ListSerialize())
		//}
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Println(">>>> serialize not implemented")
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

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
// @Param requestBody body categorydto.CreateCategoryReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /category [post]
// @Security Bearer
func CreateCategory(service categoryservice.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody categorydto.CreateCategoryReq

		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		createdCategory, err := service.CreateCategory(&requestBody, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Println(">>>> serialize not implemented")
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

		jsonResponse := presenter.NewJsonResponse(false, "", createdCategory)
		return c.Status(fiber.StatusOK).JSON(jsonResponse)
	}
}

// GetCategory
// @Summary GetCategory
// @Description GetCategory
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "category id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /category/{id} [get]
// @Security Bearer
func GetCategory(service categoryservice.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		category, err := service.GetCategoryById(id, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		//categorySerializer := serializer.NewCategorySerializer(category)
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Println(">>>> serialize not implemented")
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
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
// @Param id path int true "category id"
// @Param requestBody body categorydto.UpdateCategoryReq true "requestBody"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /category/{id} [put]
// @Security Bearer
func UpdateCategory(service categoryservice.CategoryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody categorydto.UpdateCategoryReq

		if err := c.BodyParser(&requestBody); err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		id, err := c.ParamsInt("id")
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		category, err := service.UpdateCategory(&requestBody, id, c)
		if err != nil {
			jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(jsonResponse)
		}

		//categorySerializer := serializer.NewCategorySerializer(fetchedCategory)
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Println(">>>> serialize not implemented")
		log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
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
// @Param id path int true "category id"
// @Success 200 {object} presenter.JsonResponse{}
// @Failure 503 {object} presenter.JsonResponse{}
// @Router /category/{id} [delete]
// @Security Bearer
func DeleteCategory(service categoryservice.CategoryService) fiber.Handler {
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
