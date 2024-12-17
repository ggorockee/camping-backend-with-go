package amenityrepository

import (
	amenitydto "camping-backend-with-go/internal/application/dto/amenity"
	amenityentity "camping-backend-with-go/internal/domain/entity/amenity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type AmenityRepository interface {
	CreateAmenity(input *amenitydto.CreateAmenityReq, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error)
	GetAmenityById(id int, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error)
	GetAmenityList(contexts ...*fiber.Ctx) (*[]amenityentity.Amenity, error)
	UpdateAmenity(input *amenitydto.UpdateAmenityReq, id int, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error)
	DeleteAmenity(id int, contexts ...*fiber.Ctx) error
}

type amenityRepository struct {
	dbConn *gorm.DB
}

func (r *amenityRepository) CreateAmenity(input *amenitydto.CreateAmenityReq, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error) {
	var amenity amenityentity.Amenity

	amenity.Name = input.Name
	amenity.Description = input.Description
	amenity.CreatedAt = time.Now()
	amenity.UpdatedAt = time.Now()

	if err := r.dbConn.Create(&amenity).Error; err != nil {
		return nil, err
	}

	return &amenity, nil
}

func (r *amenityRepository) GetAmenityById(id int, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error) {
	var amenity amenityentity.Amenity
	if err := r.dbConn.First(&amenity, id).Error; err != nil {
		return nil, err
	}

	return &amenity, nil
}

func (r *amenityRepository) GetAmenityList(contexts ...*fiber.Ctx) (*[]amenityentity.Amenity, error) {
	var amenities []amenityentity.Amenity
	if err := r.dbConn.Find(&amenities).Error; err != nil {
		return nil, err
	}

	return &amenities, nil
}

func (r *amenityRepository) UpdateAmenity(input *amenitydto.UpdateAmenityReq, id int, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error) {
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

func (r *amenityRepository) DeleteAmenity(id int, contexts ...*fiber.Ctx) error {
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
