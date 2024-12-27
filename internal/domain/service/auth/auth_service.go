package authservice

import (
	"camping-backend-with-go/internal/application/dto"
	authrepository "camping-backend-with-go/internal/domain/repository/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	CreateUser(input *dto.SignUpReq) error
	Login(input *dto.LoginReq) (string, error)
	ChangePassword(input *dto.ChangePasswordReq, context ...*fiber.Ctx) error
}

type authService struct {
	authRepo authrepository.AuthRepository
}

func (a *authService) CreateUser(input *dto.SignUpReq) error {
	return a.authRepo.CreateUser(input)
}

func (a *authService) Login(input *dto.LoginReq) (string, error) {
	return a.authRepo.Login(input)
}

func (a *authService) ChangePassword(input *dto.ChangePasswordReq, context ...*fiber.Ctx) error {
	return a.authRepo.ChangePassword(input, context...)
}

func NewAuthService(a authrepository.AuthRepository) AuthService {
	return &authService{authRepo: a}
}
