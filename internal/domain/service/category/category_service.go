package categoryservice

import (
	categorydto "camping-backend-with-go/internal/application/dto/category"
	categoryentity "camping-backend-with-go/internal/domain/entity/category"
	categoryrepository "camping-backend-with-go/internal/domain/repository/category"
	"github.com/gofiber/fiber/v2"
)

type CategoryService interface {
	GetCategoryList(contexts ...*fiber.Ctx) (*[]categoryentity.Category, error)
	CreateCategory(input *categorydto.CreateCategoryReq, contexts ...*fiber.Ctx) (*categoryentity.Category, error)
	UpdateCategory(input *categorydto.UpdateCategoryReq, id int, contexts ...*fiber.Ctx) (*categoryentity.Category, error)
	DeleteCategory(id int, contexts ...*fiber.Ctx) error
	GetCategoryById(id int, contexts ...*fiber.Ctx) (*categoryentity.Category, error)
}

type categoryService struct {
	categoryRepo categoryrepository.CategoryRepository
}

func (s *categoryService) GetCategoryList(contexts ...*fiber.Ctx) (*[]categoryentity.Category, error) {
	return s.categoryRepo.GetCategoryList(contexts...)
}

func (s *categoryService) CreateCategory(input *categorydto.CreateCategoryReq, contexts ...*fiber.Ctx) (*categoryentity.Category, error) {
	return s.categoryRepo.CreateCategory(input, contexts...)
}

func (s *categoryService) UpdateCategory(input *categorydto.UpdateCategoryReq, id int, contexts ...*fiber.Ctx) (*categoryentity.Category, error) {
	return s.categoryRepo.UpdateCategory(input, id, contexts...)
}

func (s *categoryService) DeleteCategory(id int, contexts ...*fiber.Ctx) error {
	return s.categoryRepo.DeleteCategory(id, contexts...)
}

func (s *categoryService) GetCategoryById(id int, contexts ...*fiber.Ctx) (*categoryentity.Category, error) {
	return s.categoryRepo.GetCategoryById(id, contexts...)
}

func NewCategoryService(c categoryrepository.CategoryRepository) CategoryService {
	return &categoryService{}
}
