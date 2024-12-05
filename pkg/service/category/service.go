package category

import (
	"camping-backend-with-go/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetCategoryList(ctx *fiber.Ctx) (*[]entities.Category, error)
	CreateCategory(createCategoryInput *entities.CreateCategoryInput, ctx *fiber.Ctx) (*entities.Category, error)
	GetCategory(id int, ctx *fiber.Ctx) (*entities.Category, error)
	UpdateCategory(updateCategoryInput *entities.UpdateCategoryInput, id int, ctx *fiber.Ctx) (*entities.Category, error)
	DeleteCategory(id int, ctx *fiber.Ctx) error
}

type service struct {
	repo Repository
}

// CreateCategory implements Service.
func (s *service) CreateCategory(createCategoryInput *entities.CreateCategoryInput, ctx *fiber.Ctx) (*entities.Category, error) {
	return s.repo.CreateCategory(createCategoryInput, ctx)
}

// DeleteCategory implements Service.
func (s *service) DeleteCategory(id int, ctx *fiber.Ctx) error {
	return s.repo.DeleteCategory(id, ctx)
}

// GetCategory implements Service.
func (s *service) GetCategory(id int, ctx *fiber.Ctx) (*entities.Category, error) {
	return s.repo.GetCategory(id, ctx)
}

// GetCategoryList implements Service.
func (s *service) GetCategoryList(ctx *fiber.Ctx) (*[]entities.Category, error) {
	return s.repo.GetCategoryList(ctx)
}

// UpdateCategory implements Service.
func (s *service) UpdateCategory(updateCategoryInput *entities.UpdateCategoryInput, id int, ctx *fiber.Ctx) (*entities.Category, error) {
	return s.repo.UpdateCategory(updateCategoryInput, id, ctx)
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
