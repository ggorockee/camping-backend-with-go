package spotservice

import (
	"camping-backend-with-go/internal/application/dto"
	"camping-backend-with-go/internal/domain/entity"
	spotrepository "camping-backend-with-go/internal/domain/repository/spot"

	"github.com/gofiber/fiber/v2"
)

type SpotService interface {
	CreateSpot(input *dto.CreateSpotReq, context ...*fiber.Ctx) (*entity.Spot, error)
	UpdateSpot(input *dto.UpdateSpotReq, id string, context ...*fiber.Ctx) (*entity.Spot, error)
	GetSpotById(id string, context ...*fiber.Ctx) (*entity.Spot, error)
	DeleteSpot(id string, context ...*fiber.Ctx) error
	GetAllSpots(context ...*fiber.Ctx) (*[]entity.Spot, error)
	GetReviewsFromSpot(spot *entity.Spot, context ...*fiber.Ctx) (*[]entity.Review, error)
	CreateSpotReview(input *dto.CreateSpotReviewReq, spot *entity.Spot, context ...*fiber.Ctx) (*entity.Review, error)
}

type spotService struct {
	spotRepo spotrepository.SpotRepository
}

func (s *spotService) CreateSpot(input *dto.CreateSpotReq, context ...*fiber.Ctx) (*entity.Spot, error) {
	return s.spotRepo.CreateSpot(input, context...)
}

func (s *spotService) UpdateSpot(input *dto.UpdateSpotReq, id string, context ...*fiber.Ctx) (*entity.Spot, error) {
	return s.spotRepo.UpdateSpot(input, id, context...)
}

func (s *spotService) GetSpotById(id string, context ...*fiber.Ctx) (*entity.Spot, error) {
	return s.spotRepo.GetSpotById(id, context...)
}

func (s *spotService) DeleteSpot(id string, context ...*fiber.Ctx) error {
	return s.spotRepo.DeleteSpot(id, context...)
}

func (s *spotService) GetAllSpots(context ...*fiber.Ctx) (*[]entity.Spot, error) {
	return s.spotRepo.GetAllSpots(context...)
}

func (s *spotService) GetReviewsFromSpot(spot *entity.Spot, context ...*fiber.Ctx) (*[]entity.Review, error) {
	return s.spotRepo.GetReviewsFromSpot(spot, context...)
}

func (s *spotService) CreateSpotReview(input *dto.CreateSpotReviewReq, spot *entity.Spot, context ...*fiber.Ctx) (*entity.Review, error) {
	return s.spotRepo.CreateSpotReview(input, spot, context...)
}

func NewSpotService(s spotrepository.SpotRepository) SpotService {
	return &spotService{spotRepo: s}
}
