package amenity

import (
	"camping-backend-with-go/internal_backup/domain"
	"camping-backend-with-go/pkg/dto"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	AddAmenity(input *dto.CreateAmenityIn, ctx *fiber.Ctx) (*entities.Amenity, error)
	GetAmenities(ctx *fiber.Ctx) (*[]entities.Amenity, error)
	GetAmenity(id int) (*entities.Amenity, error)
	UpdateAmenity(input *dto.UpdateAmenityIn, id int, ctx *fiber.Ctx) (*entities.Amenity, error)
	DeleteAmenity(id int, ctx *fiber.Ctx) error
}

type service struct {
	Repo Repository
}

func (s *service) UpdateAmenity(input *dto.UpdateAmenityIn, id int, ctx *fiber.Ctx) (*entities.Amenity, error) {
	return s.Repo.UpdateAmenity(input, id, ctx)
}

func (s *service) DeleteAmenity(id int, ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetAmenities(ctx *fiber.Ctx) (*[]entities.Amenity, error) {
	return s.Repo.GetAmenityList(ctx)
}

func (s *service) GetAmenity(id int) (*entities.Amenity, error) {
	return s.Repo.GetAmenityById(id)
}

func (s *service) AddAmenity(input *dto.CreateAmenityIn, ctx *fiber.Ctx) (*entities.Amenity, error) {
	return s.Repo.Create(input, ctx)
}

func NewService(r Repository) Service {
	return &service{
		Repo: r,
	}
}
