package categoryrepository

import (
	categorydto "camping-backend-with-go/internal/application/dto/category"
	categoryentity "camping-backend-with-go/internal/domain/entity/category"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type CategoryRepository interface {
	GetCategoryList(contexts ...*fiber.Ctx) (*[]categoryentity.Category, error)
	CreateCategory(input *categorydto.CreateCategoryReq, contexts ...*fiber.Ctx) (*categoryentity.Category, error)
	UpdateCategory(input *categorydto.UpdateCategoryReq, id int, contexts ...*fiber.Ctx) (*categoryentity.Category, error)
	DeleteCategory(id int, contexts ...*fiber.Ctx) error
	GetCategoryById(id int, contexts ...*fiber.Ctx) (*categoryentity.Category, error)
}

type categoryRepository struct {
	dbConn *gorm.DB
}

func (r *categoryRepository) GetCategoryList(contexts ...*fiber.Ctx) (*[]categoryentity.Category, error) {
	var categories []categoryentity.Category
	if err := r.dbConn.Find(&categories).Error; err != nil {
		return nil, err
	}

	return &categories, nil
}

func (r *categoryRepository) CreateCategory(input *categorydto.CreateCategoryReq, contexts ...*fiber.Ctx) (*categoryentity.Category, error) {
	var category categoryentity.Category
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	category.Name = input.Name
	if err := r.dbConn.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) GetCategoryById(id int, contexts ...*fiber.Ctx) (*categoryentity.Category, error) {
	var category categoryentity.Category
	if err := r.dbConn.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) UpdateCategory(input *categorydto.UpdateCategoryReq, id int, contexts ...*fiber.Ctx) (*categoryentity.Category, error) {
	category, err := r.GetCategoryById(id)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		category.Name = input.Name
	}

	category.UpdatedAt = time.Now()
	if err := r.dbConn.Model(&category).Updates(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) DeleteCategory(id int, contexts ...*fiber.Ctx) error {
	var category categoryentity.Category
	if err := r.dbConn.Delete(&category, id).Error; err != nil {
		return err
	}
	return nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{dbConn: db}
}
