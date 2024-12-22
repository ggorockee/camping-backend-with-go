package amenityrepository

import (
	amenitydto "camping-backend-with-go/internal/application/dto/amenity"
	"camping-backend-with-go/internal/domain/entity"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type AmenityRepository interface {
	CreateAmenity(input *amenitydto.CreateAmenityReq, context ...*fiber.Ctx) (*entity.Amenity, error)
	GetAmenityById(id string, context ...*fiber.Ctx) (*entity.Amenity, error)
	GetAmenityByName(name string, context ...*fiber.Ctx) (*entity.Amenity, error)
	GetAmenityList(context ...*fiber.Ctx) (*[]entity.Amenity, error)
	UpdateAmenity(input *amenitydto.UpdateAmenityReq, id string, context ...*fiber.Ctx) (*entity.Amenity, error)
	DeleteAmenity(id string, context ...*fiber.Ctx) error
}

type amenityRepository struct {
	dbConn *gorm.DB
}

func (r *amenityRepository) CreateAmenity(input *amenitydto.CreateAmenityReq, context ...*fiber.Ctx) (*entity.Amenity, error) {
	var amenity entity.Amenity

	if err := copier.Copy(&amenity, input); err != nil {
		return nil, err
	}

	fetched, err := r.GetAmenityByName(*input.Name)
	if err != nil {
		return nil, fmt.Errorf("error fetching amenity: %w", err)
	}

	if fetched != nil && fetched.IsExist() {
		return nil, fmt.Errorf("amenity name is duplicated")
	}

	err = r.dbConn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&amenity).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &amenity, nil
}

func (r *amenityRepository) GetAmenityById(id string, context ...*fiber.Ctx) (*entity.Amenity, error) {
	var amenity entity.Amenity

	err := r.dbConn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).First(&amenity).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &amenity, nil
}

func (r *amenityRepository) GetAmenityByName(name string, context ...*fiber.Ctx) (*entity.Amenity, error) {
	var amenity entity.Amenity

	err := r.dbConn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("name = ?", name).First(&amenity).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
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

func (r *amenityRepository) UpdateAmenity(input *amenitydto.UpdateAmenityReq, id string, context ...*fiber.Ctx) (*entity.Amenity, error) {
	amenity, err := r.GetAmenityById(id)
	if err != nil {
		return nil, err
	}

	if err := copier.Copy(amenity, input); err != nil {
		return nil, err
	}

	if err := r.dbConn.Save(amenity).Error; err != nil {
		return nil, err
	}

	return amenity, nil
}

func (r *amenityRepository) DeleteAmenity(id string, context ...*fiber.Ctx) error {
	amenity, err := r.GetAmenityById(id)
	if err != nil {
		return err
	}

	if err := r.dbConn.Where("id = ?", id).Delete(amenity).Error; err != nil {
		return err
	}

	return nil
}

func NewAmenityRepository(db *gorm.DB) AmenityRepository {
	return &amenityRepository{dbConn: db}
}
