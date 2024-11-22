package spot

import (
	"camping-backend-with-go/pkg/entities"
	"gorm.io/gorm"
)

type SpotRepository interface {
	CreateSpot(spot *entities.Spot) (*entities.Spot, error)
	//ReadSpot() (*[]presenter.Spot, error)
	//UpdateSpot(spot *entities.Spot) (*entities.Spot, error)
	//DeleteSpot(Id string) error
}

type repository struct {
	DBConn *gorm.DB
}

func NewRepo(dbconn *gorm.DB) SpotRepository {
	return &repository{
		DBConn: dbconn,
	}
}

func (r *repository) CreateSpot(spot *entities.Spot) (*entities.Spot, error) {
	result := r.DBConn.Create(spot)
	if result.Error != nil {
		return nil, result.Error
	}
	return spot, nil

}

//func (r *repository) ReadSpot() (*[]presenter.Spot, error) {}
//
//func (r *repository) UpdateSpot(spot *entities.Spot) (*entities.Spot, error) {}
//
//func (r *repository) DeleteSpot(Id string) error {}
