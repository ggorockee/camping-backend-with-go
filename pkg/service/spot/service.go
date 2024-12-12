package spot

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	InsertSpot(input *dto.CreateSpotIn, ctx *fiber.Ctx) (*entities.Spot, error)
	UpdateSpot(input *dto.UpdateSpotIn, id int, ctx *fiber.Ctx) (*entities.Spot, error)
	GetSpot(id int, ctx ...*fiber.Ctx) (*entities.Spot, error)
	RemoveSpot(id int, ctx *fiber.Ctx) error
	GetAllSpots() (*[]entities.Spot, error)
	GetReviewsFromSpot(spot *entities.Spot, contexts ...*fiber.Ctx) (*[]entities.Review, error)
	CreateSpotReview(input *dto.CreateSpotReviewReq, spot *entities.Spot, contexts ...*fiber.Ctx) (*entities.Review, error)
}

type service struct {
	repository Repository
}

// CreateSpotReview implements Service.
func (s *service) CreateSpotReview(input *dto.CreateSpotReviewReq, spot *entities.Spot, contexts ...*fiber.Ctx) (*entities.Review, error) {
	return s.repository.CreateSpotReview(input, spot, contexts...)
}

func (s *service) GetReviewsFromSpot(spot *entities.Spot, contexts ...*fiber.Ctx) (*[]entities.Review, error) {
	if len(contexts) > 0 {
		return s.repository.GetReviewsFromSpot(spot, contexts[0])
	}
	return s.repository.GetReviewsFromSpot(spot)
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
func (s *service) InsertSpot(input *dto.CreateSpotIn, ctx *fiber.Ctx) (*entities.Spot, error) {
	return s.repository.CreateSpot(input, ctx)
}

// UpdateSpot is a service layer that helps update Spots in Camping
func (s *service) UpdateSpot(input *dto.UpdateSpotIn, id int, ctx *fiber.Ctx) (*entities.Spot, error) {
	return s.repository.UpdateSpot(input, id, ctx)
}

// GetSpot is a service layer that helps update Spots in Camping
func (s *service) GetSpot(id int, ctx ...*fiber.Ctx) (*entities.Spot, error) {
	return s.repository.GetSpot(id, ctx...)
}

func (s *service) RemoveSpot(id int, ctx *fiber.Ctx) error {

	return s.repository.DeleteSpot(id, ctx)
}
