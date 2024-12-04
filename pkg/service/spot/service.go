package spot

import (
	"camping-backend-with-go/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	InsertSpot(createSpotInputSchema *entities.CreateSpotInputSchema, ctx *fiber.Ctx) (*entities.Spot, error)
	FetchMySpots(ctx *fiber.Ctx) (*[]entities.Spot, error)
	UpdateSpot(spot *entities.UpdateSpotSchema, id int, ctx *fiber.Ctx) (*entities.Spot, error)
	GetSpot(id int, ctx *fiber.Ctx) (*entities.Spot, error)
	RemoveSpot(id int, ctx *fiber.Ctx) error
	GetAllSpots() (*[]entities.Spot, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// GetAllSpots implements Service.
func (s *service) GetAllSpots() (*[]entities.Spot, error) {
	return s.repository.GetAllSpots()
}

// InsertSpot is a service layer that helps insert Spot in Camping
func (s *service) InsertSpot(createSpotInputSchema *entities.CreateSpotInputSchema, ctx *fiber.Ctx) (*entities.Spot, error) {
	return s.repository.CreateSpot(createSpotInputSchema, ctx)
}

// FetchSpots is a service layer that helps fetch all Spots in Camping
func (s *service) FetchMySpots(ctx *fiber.Ctx) (*[]entities.Spot, error) {
	return s.repository.FetchMySpots(ctx)
}

// UpdateSpot is a service layer that helps update Spots in Camping
func (s *service) UpdateSpot(spot *entities.UpdateSpotSchema, id int, ctx *fiber.Ctx) (*entities.Spot, error) {
	return s.repository.UpdateSpot(spot, id, ctx)
}

// GetSpot is a service layer that helps update Spots in Camping
func (s *service) GetSpot(id int, ctx *fiber.Ctx) (*entities.Spot, error) {
	return s.repository.GetSpot(id, ctx)
}

func (s *service) RemoveSpot(id int, ctx *fiber.Ctx) error {
	return s.repository.DeleteSpot(id, ctx)
}
