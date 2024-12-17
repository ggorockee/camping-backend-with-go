package amenityservice

import (
	amenitydto "camping-backend-with-go/internal/application/dto/amenity"
	amenityentity "camping-backend-with-go/internal/domain/entity/amenity"
	amenityrepository "camping-backend-with-go/internal/domain/repository/amenity"
	"github.com/gofiber/fiber/v2"
)

type AmenityService interface {
	CreateAmenity(input *amenitydto.CreateAmenityReq, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error)
	GetAmenityById(id int, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error)
	GetAmenityList(contexts ...*fiber.Ctx) (*[]amenityentity.Amenity, error)
	UpdateAmenity(input *amenitydto.UpdateAmenityReq, id int, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error)
	DeleteAmenity(id int, contexts ...*fiber.Ctx) error
}

type amenityService struct {
	amenService amenityrepository.AmenityRepository
}

func (s *amenityService) CreateAmenity(input *amenitydto.CreateAmenityReq, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error) {
	return s.amenService.CreateAmenity(input, contexts...)
}

func (s *amenityService) GetAmenityById(id int, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error) {
	return s.amenService.GetAmenityById(id, contexts...)
}

func (s *amenityService) GetAmenityList(contexts ...*fiber.Ctx) (*[]amenityentity.Amenity, error) {
	return s.amenService.GetAmenityList(contexts...)
}

func (s *amenityService) UpdateAmenity(input *amenitydto.UpdateAmenityReq, id int, contexts ...*fiber.Ctx) (*amenityentity.Amenity, error) {
	return s.amenService.UpdateAmenity(input, id, contexts...)
}

func (s *amenityService) DeleteAmenity(id int, contexts ...*fiber.Ctx) error {
	return s.amenService.DeleteAmenity(id, contexts...)
}

func NewAmenityService(r amenityrepository.AmenityRepository) AmenityService {
	return &amenityService{amenService: r}
}
