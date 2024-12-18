package amenityrepository

import (
	amenitydto "camping-backend-with-go/internal/application/dto/amenity"
	"camping-backend-with-go/internal/domain/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type AmenityRepository interface {
	CreateAmenity(input *amenitydto.CreateAmenityReq, context ...*fiber.Ctx) (*entity.Amenity, error)
	GetAmenityById(id int, context ...*fiber.Ctx) (*entity.Amenity, error)
	GetAmenityList(context ...*fiber.Ctx) (*[]entity.Amenity, error)
	UpdateAmenity(input *amenitydto.UpdateAmenityReq, id int, context ...*fiber.Ctx) (*entity.Amenity, error)
	DeleteAmenity(id int, context ...*fiber.Ctx) error
}

type amenityRepository struct {
	dbConn *gorm.DB
}

func (r *amenityRepository) CreateAmenity(input *amenitydto.CreateAmenityReq, context ...*fiber.Ctx) (*entity.Amenity, error) {
	var amenity entity.Amenity

	amenity.Name = input.Name
	amenity.Description = input.Description
	amenity.CreatedAt = time.Now()
	amenity.UpdatedAt = time.Now()

	if err := r.dbConn.Create(&amenity).Error; err != nil {
		return nil, err
	}

	return &amenity, nil
}

func (r *amenityRepository) GetAmenityById(id int, context ...*fiber.Ctx) (*entity.Amenity, error) {
	var amenity entity.Amenity
	if err := r.dbConn.First(&amenity, id).Error; err != nil {
		return nil, err
	}

	return &amenity, nil
}

func (r *amenityRepository) GetAmenityList(context ...*fiber.Ctx) (*[]entity.Amenity, error) {
	var amenities []entity.Amenity
	if err := r.dbConn.Find(&amenities).Error; err != nil {
		return nil, err
	}

	return &amenities, nil
}

func (r *amenityRepository) UpdateAmenity(input *amenitydto.UpdateAmenityReq, id int, context ...*fiber.Ctx) (*entity.Amenity, error) {
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

	if err := r.dbConn.Model(amenity).Updates(amenity).Error; err != nil {
		return nil, err
	}

	return amenity, nil
}

func (r *amenityRepository) DeleteAmenity(id int, context ...*fiber.Ctx) error {
	amenity, err := r.GetAmenityById(id)
	if err != nil {
		return err
	}

	if err := r.dbConn.Delete(&amenity).Error; err != nil {
		return err
	}

	return nil
}

func NewAmenityRepository(db *gorm.DB) AmenityRepository {
	return &amenityRepository{dbConn: db}
}
