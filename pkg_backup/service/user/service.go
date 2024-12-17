package user

import (
	"camping-backend-with-go/pkg/dto"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	CreateUser(input *dto.SignUpIn) error
	Login(input *dto.LoginIn) (string, error)
	ChangePassword(input *dto.ChangePasswordIn, ctx *fiber.Ctx) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Login(input *dto.LoginIn) (string, error) {
	return s.repository.Login(input)
}

func (s *service) CreateUser(input *dto.SignUpIn) error {
	return s.repository.CreateUser(input)
}

func (s *service) ChangePassword(input *dto.ChangePasswordIn, ctx *fiber.Ctx) error {
	return s.repository.ChangePassword(input, ctx)
}
