package categoryservice

import (
	categorydto "camping-backend-with-go/internal/application/dto/category"
	"camping-backend-with-go/internal/domain/entity"
	categoryrepository "camping-backend-with-go/internal/domain/repository/category"

	"github.com/gofiber/fiber/v2"
)

type CategoryService interface {
	GetCategoryList(context ...*fiber.Ctx) (*[]entity.Category, error)
	CreateCategory(input *categorydto.CreateCategoryReq, context ...*fiber.Ctx) (*entity.Category, error)
	UpdateCategory(input *categorydto.UpdateCategoryReq, id string, context ...*fiber.Ctx) (*entity.Category, error)
	DeleteCategory(id string, context ...*fiber.Ctx) error
	GetCategoryById(id string, context ...*fiber.Ctx) (*entity.Category, error)
}

type categoryService struct {
	categoryRepo categoryrepository.CategoryRepository
}

func (s *categoryService) GetCategoryList(context ...*fiber.Ctx) (*[]entity.Category, error) {
	return s.categoryRepo.GetCategoryList(context...)
}

func (s *categoryService) CreateCategory(input *categorydto.CreateCategoryReq, context ...*fiber.Ctx) (*entity.Category, error) {
	return s.categoryRepo.CreateCategory(input, context...)
}

func (s *categoryService) UpdateCategory(input *categorydto.UpdateCategoryReq, id string, context ...*fiber.Ctx) (*entity.Category, error) {
	return s.categoryRepo.UpdateCategory(input, id, context...)
}

func (s *categoryService) DeleteCategory(id string, context ...*fiber.Ctx) error {
	return s.categoryRepo.DeleteCategory(id, context...)
}

func (s *categoryService) GetCategoryById(id string, context ...*fiber.Ctx) (*entity.Category, error) {
	return s.categoryRepo.GetCategoryById(id, context...)
}

func NewCategoryService(c categoryrepository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: c}
}
