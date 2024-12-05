package handlers

import (
	"camping-backend-with-go/pkg/service/category"

	"github.com/gofiber/fiber/v2"
)

func GetCategoryList(service category.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
func CreateCategory(service category.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
func GetCategory(service category.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
func UpdateCategory(service category.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
func DeleteCategory(service category.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
