package spot

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
)

type Service interface {
	InsertSpot(spot *entities.Spot) (*entities.Spot, error)
	FetchSpots() (*[]presenter.Spot, error)
	UpdateSpot(spot *entities.Spot) (*entities.Spot, error)
	RemoveSpot(ID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// InsertSpot is a service layer that helps insert Spot in SpotShop
func (s *service) InsertSpot(spot *entities.Spot) (*entities.Spot, error) {
	return s.repository.CreateSpot(spot)
}

// FetchSpots is a service layer that helps fetch all Spots in SpotShop
func (s *service) FetchSpots() (*[]presenter.Spot, error) {
	return s.repository.ReadSpot()
}

// UpdateSpot is a service layer that helps update Spots in SpotShop
func (s *service) UpdateSpot(spot *entities.Spot) (*entities.Spot, error) {
	return s.repository.UpdateSpot(spot)
}

// RemoveSpot is a service layer that helps remove Spots from SpotShop
func (s *service) RemoveSpot(ID string) error {
	return s.repository.DeleteSpot(ID)
}
