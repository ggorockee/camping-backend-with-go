package user

import (
	"camping-backend-with-go/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *entities.User) (*entities.User, error)
}

type repository struct {
	DBConn *gorm.DB
}

func NewRepo(dbconn *gorm.DB) Repository {
	return &repository{
		DBConn: dbconn,
	}
}

func (r *repository) CreateUser(user *entities.User) (*entities.User, error) {
	result := r.DBConn.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
