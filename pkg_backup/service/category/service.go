package category

import (
	"camping-backend-with-go/internal_backup/domain"
	"camping-backend-with-go/pkg/dto"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetCategoryList(ctx *fiber.Ctx) (*[]entities.Category, error)
	CreateCategory(input *dto.CreateCategoryIn, ctx *fiber.Ctx) (*entities.Category, error)
	GetCategory(id int, ctx *fiber.Ctx) (*entities.Category, error)
	UpdateCategory(input *dto.UpdateCategoryIn, id int, ctx *fiber.Ctx) (*entities.Category, error)
	DeleteCategory(id int, ctx *fiber.Ctx) error
}

type service struct {
	repo Repository
}

// CreateCategory implements Service.
func (s *service) CreateCategory(input *dto.CreateCategoryIn, ctx *fiber.Ctx) (*entities.Category, error) {
	return s.repo.CreateCategory(input, ctx)
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
func (s *service) UpdateCategory(input *dto.UpdateCategoryIn, id int, ctx *fiber.Ctx) (*entities.Category, error) {
	return s.repo.UpdateCategory(input, id, ctx)
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
