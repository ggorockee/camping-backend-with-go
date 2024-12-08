package amenity

import (
	"camping-backend-with-go/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	AddAmenity(input *entities.CreateAmenityInput, ctx *fiber.Ctx) (*entities.Amenity, error)
	GetAmenities(ctx *fiber.Ctx) (*[]entities.Amenity, error)
	GetAmenity(id int, ctx *fiber.Ctx) (*entities.Amenity, error)
	UpdateAmenity(input *entities.UpdateAmenityInput, id int, ctx *fiber.Ctx) (*entities.Amenity, error)
	DeleteAmenity(id int, ctx *fiber.Ctx) error
}

type service struct {
	Repo Repository
}

func (s *service) UpdateAmenity(input *entities.UpdateAmenityInput, id int, ctx *fiber.Ctx) (*entities.Amenity, error) {
	return s.Repo.UpdateAmenity(input, id, ctx)
}

func (s *service) DeleteAmenity(id int, ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetAmenities(ctx *fiber.Ctx) (*[]entities.Amenity, error) {
	return s.Repo.GetAmenityList(ctx)
}

func (s *service) GetAmenity(id int, ctx *fiber.Ctx) (*entities.Amenity, error) {
	return s.Repo.GetAmenityById(id, ctx)
}

func (s *service) AddAmenity(input *entities.CreateAmenityInput, ctx *fiber.Ctx) (*entities.Amenity, error) {
	return s.Repo.Create(input, ctx)
}

func NewService(r Repository) Service {
	return &service{
		Repo: r,
	}
}
