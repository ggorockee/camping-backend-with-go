package userservice

import (
	"camping-backend-with-go/internal/application/dto"
	"camping-backend-with-go/internal/domain/entity"
	userrepository "camping-backend-with-go/internal/domain/repository/user"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	HashPassword(password string, context ...*fiber.Ctx) (string, error)
	GetUserByEmail(email string, context ...*fiber.Ctx) (*entity.User, error)
	CheckPasswordHash(password, hash string, context ...*fiber.Ctx) bool      // auth
	ChangePassword(input *dto.ChangePasswordReq, context ...*fiber.Ctx) error //
	ValidToken(t *jwt.Token, id string, context ...*fiber.Ctx) bool
	GetUserById(id string, context ...*fiber.Ctx) (*entity.User, error)
	GetValueFromToken(key string, context ...*fiber.Ctx) string
}

type userService struct {
	userRepo userrepository.UserRepository
}

func (s *userService) HashPassword(password string, context ...*fiber.Ctx) (string, error) {
	return s.userRepo.HashPassword(password, context...)
}

func (s *userService) GetUserByEmail(email string, context ...*fiber.Ctx) (*entity.User, error) {
	return s.userRepo.GetUserByEmail(email, context...)
}

func (s *userService) CheckPasswordHash(password, hash string, context ...*fiber.Ctx) bool {
	return s.userRepo.CheckPasswordHash(password, hash, context...)
}

func (s *userService) ChangePassword(input *dto.ChangePasswordReq, context ...*fiber.Ctx) error {
	return s.userRepo.ChangePassword(input, context...)
}

func (s *userService) ValidToken(t *jwt.Token, id string, context ...*fiber.Ctx) bool {
	return s.userRepo.ValidToken(t, id, context...)
}

func (s *userService) GetUserById(id string, context ...*fiber.Ctx) (*entity.User, error) {
	return s.userRepo.GetUserById(id, context...)
}

func (s *userService) GetValueFromToken(key string, context ...*fiber.Ctx) string {
	return s.userRepo.GetValueFromToken(key, context...)
}

func NewUserService(u userrepository.UserRepository) UserService {
	return &userService{userRepo: u}
}
