package amenity

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	Create(input *dto.CreateAmenityIn, ctx *fiber.Ctx) (*entities.Amenity, error)
	GetAmenityById(id int) (*entities.Amenity, error)
	GetAmenityList(ctx *fiber.Ctx) (*[]entities.Amenity, error)
	UpdateAmenity(input *dto.UpdateAmenityIn, id int, ctx *fiber.Ctx) (*entities.Amenity, error)
	DeleteAmenity(id int, ctx *fiber.Ctx) error
}

type repository struct {
	DBConn   *gorm.DB
	UserRepo user.Repository
}

func (r *repository) DeleteAmenity(id int, ctx *fiber.Ctx) error {
	amenity, err := r.GetAmenityById(id)
	if err != nil {
		return err
	}

	if err := r.DBConn.Delete(&amenity).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateAmenity(input *dto.UpdateAmenityIn, id int, ctx *fiber.Ctx) (*entities.Amenity, error) {
	amenity, err := r.GetAmenityById(id)
	if err != nil {
		return nil, err
	}

	amenity.UpdatedAt = time.Now()
	switch {
	case input.Name != "":
		amenity.Name = input.Name
	case input.Description != nil:
		amenity.Description = input.Description
	}

	var model entities.Amenity
	model.Id = amenity.Id
	if err := r.DBConn.Model(model).Updates(&amenity).Error; err != nil {
		return nil, err
	}

	return amenity, nil
}

func (r *repository) GetAmenityById(id int) (*entities.Amenity, error) {
	var amenity entities.Amenity
	if err := r.DBConn.First(&amenity, id).Error; err != nil {
		return nil, err
	}

	return &amenity, nil
}

func (r *repository) GetAmenityList(ctx *fiber.Ctx) (*[]entities.Amenity, error) {
	var amenities []entities.Amenity
	if err := r.DBConn.Find(&amenities).Error; err != nil {
		return nil, err
	}

	return &amenities, nil
}

func (r *repository) Create(input *dto.CreateAmenityIn, ctx *fiber.Ctx) (*entities.Amenity, error) {
	_ = r.UserRepo.GetValueFromToken("user_id", ctx)

	var amenity entities.Amenity

	amenity.Name = input.Name
	amenity.Description = input.Description
	amenity.CreatedAt = time.Now()
	amenity.UpdatedAt = time.Now()

	if err := r.DBConn.Create(&amenity).Error; err != nil {
		return nil, err
	}

	return &amenity, nil
}

func NewRepo(dbConn *gorm.DB, userRepo user.Repository) Repository {
	return &repository{DBConn: dbConn, UserRepo: userRepo}
}
