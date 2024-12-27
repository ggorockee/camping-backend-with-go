package amenityservice

import (
	"camping-backend-with-go/internal/application/dto"
	"camping-backend-with-go/internal/domain/entity"
	amenityrepository "camping-backend-with-go/internal/domain/repository/amenity"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type AmenityService interface {
	CreateAmenity(input *dto.CreateAmenityReq, context ...*fiber.Ctx) (*entity.Amenity, error)
	GetAmenityById(id string, context ...*fiber.Ctx) (*entity.Amenity, error)
	GetAmenityList(context ...*fiber.Ctx) (*[]entity.Amenity, error)
	UpdateAmenity(input *dto.UpdateAmenityReq, id string, context ...*fiber.Ctx) (*entity.Amenity, error)
	DeleteAmenity(id string, context ...*fiber.Ctx) error
}

type amenityService struct {
	amenService amenityrepository.AmenityRepository
}

func (s *amenityService) CreateAmenity(input *dto.CreateAmenityReq, context ...*fiber.Ctx) (*entity.Amenity, error) {
	// 이름 길이 검증
	if len(*input.Name) > 10 {
		return nil, errors.New("name is too long")
	}
	return s.amenService.CreateAmenity(input, context...)
}

func (s *amenityService) GetAmenityById(id string, context ...*fiber.Ctx) (*entity.Amenity, error) {
	return s.amenService.GetAmenityById(id, context...)
}

func (s *amenityService) GetAmenityList(context ...*fiber.Ctx) (*[]entity.Amenity, error) {
	return s.amenService.GetAmenityList(context...)
}

func (s *amenityService) UpdateAmenity(input *dto.UpdateAmenityReq, id string, context ...*fiber.Ctx) (*entity.Amenity, error) {
	return s.amenService.UpdateAmenity(input, id, context...)
}

func (s *amenityService) DeleteAmenity(id string, context ...*fiber.Ctx) error {
	return s.amenService.DeleteAmenity(id, context...)
}

func NewAmenityService(r amenityrepository.AmenityRepository) AmenityService {
	return &amenityService{amenService: r}
}
