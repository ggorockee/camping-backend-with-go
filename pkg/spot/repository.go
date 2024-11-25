package spot

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	CreateSpot(spot *entities.Spot) (*entities.Spot, error)
	ReadSpot() (*[]presenter.Spot, error)
	GetSpot(id int) (*entities.Spot, error)
	UpdateSpot(spot *entities.Spot, id int) (*entities.Spot, error)
	GetFindById(id int) (*entities.Spot, error)
	PartialUpdateSpot(spot *entities.Spot, id int) (*entities.Spot, error)
	DeleteSpot(id int) error
}

type repository struct {
	DBConn *gorm.DB
}

func NewRepo(dbconn *gorm.DB) Repository {
	return &repository{
		DBConn: dbconn,
	}
}

func (r *repository) GetSpot(id int) (*entities.Spot, error) {
	var spot entities.Spot
	result := r.DBConn.First(&spot, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &spot, nil
}

func (r *repository) GetFindById(id int) (*entities.Spot, error) {
	var spot entities.Spot
	result := r.DBConn.First(&spot, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &spot, nil
}

func (r *repository) CreateSpot(spot *entities.Spot) (*entities.Spot, error) {
	result := r.DBConn.Create(spot)
	if result.Error != nil {
		return nil, result.Error
	}
	return spot, nil

}

func (r *repository) ReadSpot() (*[]presenter.Spot, error) {
	var spots []presenter.Spot
	result := r.DBConn.Find(&spots)
	if result.Error != nil {
		return nil, result.Error
	}

	return &spots, nil
}

func (r *repository) UpdateSpot(spot *entities.Spot, id int) (*entities.Spot, error) {
	fetched, err := r.GetFindById(id)

	if err != nil {
		return nil, err
	}

	spot.UpdatedAt = time.Now()

	result := r.DBConn.Model(fetched).Updates(&spot)

	if result.Error != nil {
		return nil, result.Error
	}

	return spot, nil
}

func (r *repository) PartialUpdateSpot(spot *entities.Spot, id int) (*entities.Spot, error) {
	fetched, err := r.GetFindById(id)
	if err != nil {
		return nil, err
	}

	spot.UpdatedAt = time.Now()

	result := r.DBConn.Model(fetched).Updates(&spot)
	if result.Error != nil {
		return nil, result.Error
	}

	return fetched, nil
}

func (r *repository) DeleteSpot(id int) error {
	spot, err := r.GetFindById(id)
	if err != nil {
		return err
	}

	result := r.DBConn.Delete(spot)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
