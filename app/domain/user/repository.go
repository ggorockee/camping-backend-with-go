package user

import (
	"camping-backend-with-go/app/core/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByEmail(email string) (*model.User, error)
}

type repository struct {
	dbConn *gorm.DB
}

// GetUserByEmail implements Repository.
func (r *repository) GetUserByEmail(email string) (*model.User, error) {
	panic("unimplemented")
}

func NewRepository(dbConn *gorm.DB) Repository {
	return &repository{
		dbConn: dbConn,
	}
}
