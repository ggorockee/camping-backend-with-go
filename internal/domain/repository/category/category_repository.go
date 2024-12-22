package categoryrepository

import (
	categorydto "camping-backend-with-go/internal/application/dto/category"
	"camping-backend-with-go/internal/domain/entity"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoryList(context ...*fiber.Ctx) (*[]entity.Category, error)
	CreateCategory(input *categorydto.CreateCategoryReq, context ...*fiber.Ctx) (*entity.Category, error)
	UpdateCategory(input *categorydto.UpdateCategoryReq, id string, context ...*fiber.Ctx) (*entity.Category, error)
	DeleteCategory(id string, context ...*fiber.Ctx) error
	GetCategoryById(id string, context ...*fiber.Ctx) (*entity.Category, error)
	GetCategoryByName(name string, context ...*fiber.Ctx) (*entity.Category, error)
}

type categoryRepository struct {
	dbConn *gorm.DB
}

func (r *categoryRepository) GetCategoryList(context ...*fiber.Ctx) (*[]entity.Category, error) {
	var categories []entity.Category
	if err := r.dbConn.Find(&categories).Error; err != nil {
		return nil, err
	}

	return &categories, nil
}

func (r *categoryRepository) CreateCategory(input *categorydto.CreateCategoryReq, context ...*fiber.Ctx) (*entity.Category, error) {
	var category entity.Category

	if err := copier.Copy(&category, input); err != nil {
		return nil, err
	}

	fetched, err := r.GetCategoryByName(*input.Name)
	if err != nil {
		return nil, fmt.Errorf("error fetching category: %w", err)
	}
	if fetched != nil && fetched.IsExist() {
		return nil, fmt.Errorf("category name is duplicated")
	}

	err = r.dbConn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&category).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) GetCategoryById(id string, context ...*fiber.Ctx) (*entity.Category, error) {
	var category entity.Category
	if err := r.dbConn.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) GetCategoryByName(name string, context ...*fiber.Ctx) (*entity.Category, error) {
	var category entity.Category

	err := r.dbConn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("name = ?", name).Find(&category).Error; err != nil {
			return fmt.Errorf("error fetching category %w", err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) UpdateCategory(input *categorydto.UpdateCategoryReq, id string, context ...*fiber.Ctx) (*entity.Category, error) {
	category, err := r.GetCategoryById(id)
	if err != nil {
		return nil, err
	}

	if err := copier.Copy(category, input); err != nil {
		return nil, err
	}

	if err := r.dbConn.Save(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) DeleteCategory(id string, context ...*fiber.Ctx) error {
	var category entity.Category
	if err := r.dbConn.Where("id = ?", id).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{dbConn: db}
}
