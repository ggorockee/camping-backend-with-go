package user

import "camping-backend-with-go/pkg/entities"

type Service interface {
	CreateUser(user *entities.User) (*entities.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateUser(user *entities.User) (*entities.User, error) {
	return s.repository.CreateUser(user)
}
