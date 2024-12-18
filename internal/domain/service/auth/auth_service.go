package authservice

import (
	authdto "camping-backend-with-go/internal/application/dto/auth"
	authrepository "camping-backend-with-go/internal/domain/repository/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	CreateUser(input *authdto.SignUpReq) error
	Login(input *authdto.LoginReq) (string, error)
	ChangePassword(input *authdto.ChangePasswordReq, context ...*fiber.Ctx) error
}

type authService struct {
	authRepo authrepository.AuthRepository
}

func (a *authService) CreateUser(input *authdto.SignUpReq) error {
	return a.authRepo.CreateUser(input)
}

func (a *authService) Login(input *authdto.LoginReq) (string, error) {
	return a.authRepo.Login(input)
}

func (a *authService) ChangePassword(input *authdto.ChangePasswordReq, context ...*fiber.Ctx) error {
	return a.authRepo.ChangePassword(input, context...)
}

func NewAuthService(a authrepository.AuthRepository) AuthService {
	return &authService{authRepo: a}
}
