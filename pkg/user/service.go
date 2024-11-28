package user

import "camping-backend-with-go/pkg/entities"

type Service interface {
	CreateUser(signUpInputSchema *entities.SignUpInputSchema) error
	Login(loginInputSchema *entities.LoginInputSchema) (string, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Login(loginInputSchema *entities.LoginInputSchema) (string, error) {
	return s.repository.Login(loginInputSchema)
}

func (s *service) CreateUser(signUpInputSchema *entities.SignUpInputSchema) error {
	return s.repository.CreateUser(signUpInputSchema)
}
