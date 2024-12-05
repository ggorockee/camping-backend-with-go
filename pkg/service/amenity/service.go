package amenity

import (
	"camping-backend-with-go/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	AddAmenity(input *entities.CreateAmenityInput, ctx *fiber.Ctx) (*entities.Amenity, error)
}

type service struct {
	Repo Repository
}

func (s *service) AddAmenity(input *entities.CreateAmenityInput, ctx *fiber.Ctx) (*entities.Amenity, error) {
	return s.Repo.Create(input, ctx)
}

func NewService(r Repository) Service {
	return &service{
		Repo: r,
	}
}
