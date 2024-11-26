package auth

import "camping-backend-with-go/pkg/entities"

type Service interface {
	Login(loginInput *entities.Login) (string, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Login(loginInput *entities.Login) (string, error) {
	return s.repository.Login(loginInput)
}
