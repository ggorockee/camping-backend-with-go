package category

import (
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/user"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository interface {
	GetCategoryList(ctx *fiber.Ctx) (*[]entities.Category, error)
	CreateCategory(createCategoryInput *entities.CreateCategoryInput, ctx *fiber.Ctx) (*entities.Category, error)
	GetCategory(id int, ctx *fiber.Ctx) (*entities.Category, error)
	UpdateCategory(updateCategoryInput *entities.UpdateCategoryInput, id int, ctx *fiber.Ctx) (*entities.Category, error)
	DeleteCategory(id int, ctx *fiber.Ctx) error
	GetCategoryById(id int) (*entities.Category, error)
}

type repository struct {
	DBConn   *gorm.DB
	UserRepo user.Repository
}

// GetCategoryById implements Repository.
func (r *repository) GetCategoryById(id int) (*entities.Category, error) {
	var category entities.Category
	if err := r.DBConn.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// CreateCategory implements Repository.
func (r *repository) CreateCategory(createCategoryInput *entities.CreateCategoryInput, ctx *fiber.Ctx) (*entities.Category, error) {
	// Login 인증
	userId := r.UserRepo.GetValueFromToken("user_id", ctx)
	LoginUser, err := r.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	// validation
	// Todo
	// admin만 가능
	log.Println("Todo: staff이거나 admin인 경우만 생성할 수 있어야함", LoginUser.Email)

	var category entities.Category

	// time log
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	category.Name = createCategoryInput.Name

	if err := r.DBConn.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// DeleteCategory implements Repository.
func (r *repository) DeleteCategory(id int, ctx *fiber.Ctx) error {
	userId := r.UserRepo.GetValueFromToken("user_id", ctx)

	// Validation
	// 1. Admin이나 staff가 아닐 경우 fail
	log.Println("Validation Admin이나 staff가 아닐 경우 fail", userId)

	var category entities.Category
	if err := r.DBConn.Delete(&category, id).Error; err != nil {
		return err
	}
	return nil
}

// GetCategory implements Repository.
func (r *repository) GetCategory(id int, ctx *fiber.Ctx) (*entities.Category, error) {
	// Login인증
	userId := r.UserRepo.GetValueFromToken("user_id", ctx)

	// Validation
	// 1. Admin이나 staff가 아닐 경우 fail
	log.Println("Validation Admin이나 staff가 아닐 경우 fail", userId)

	var category entities.Category
	if err := r.DBConn.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// GetCategoryList implements Repository.
func (r *repository) GetCategoryList(ctx *fiber.Ctx) (*[]entities.Category, error) {
	// Login인증
	userId := r.UserRepo.GetValueFromToken("user_id", ctx)
	_, err := r.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	var categories []entities.Category

	// 필요하면 orderby 추가하기로..
	if err := r.DBConn.Find(&categories).Error; err != nil {
		return nil, err
	}

	return &categories, nil
}

// UpdateCategory implements Repository.
func (r *repository) UpdateCategory(updateCategoryInput *entities.UpdateCategoryInput, id int, ctx *fiber.Ctx) (*entities.Category, error) {
	// Login 인증
	userId := r.UserRepo.GetValueFromToken("user_id", ctx)

	// Get Category
	fetchedCategory, err := r.GetCategoryById(id)
	if err != nil {
		return nil, err
	}

	// Validation
	// 1. Admin이나 staff가 아닐 경우 fail
	log.Println("Validation Admin이나 staff가 아닐 경우 fail", userId)

	if updateCategoryInput.Name != "" {
		fetchedCategory.Name = updateCategoryInput.Name
	}

	fetchedCategory.UpdatedAt = time.Now()

	var category entities.Category
	category.Id = fetchedCategory.Id
	if err := r.DBConn.Model(&category).Updates(fetchedCategory).Error; err != nil {
		return nil, err
	}

	return fetchedCategory, nil
}

func NewRepo(dbconn *gorm.DB, userRepo user.Repository) Repository {
	return &repository{
		DBConn:   dbconn,
		UserRepo: userRepo,
	}
}
