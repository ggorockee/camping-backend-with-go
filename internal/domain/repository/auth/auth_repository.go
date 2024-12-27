package authrepository

import (
	"camping-backend-with-go/internal/application/dto"
	"camping-backend-with-go/internal/domain/entity"
	"camping-backend-with-go/pkg/config"
	"camping-backend-with-go/pkg/util"
	"errors"
	"strings"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(input *dto.LoginReq) (string, error)
	CreateUser(input *dto.SignUpReq) error
	ChangePassword(input *dto.ChangePasswordReq, context ...*fiber.Ctx) error

	hashPassword(password string) (string, error)
	GetUserByEmail(email string) (*entity.User, error)
	CheckPasswordHash(password, hash string) bool
	ValidToken(t *jwt.Token, id string) bool
	GetUserById(id int) (*entity.User, error)
	GetValueFromToken(key string, c *fiber.Ctx) int
}

type authRepository struct {
	dbConn *gorm.DB
}

func (a *authRepository) Login(input *dto.LoginReq) (string, error) {
	email := input.Email
	password := input.Password

	user, err := a.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !a.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (a *authRepository) CreateUser(input *dto.SignUpReq) error {
	var user entity.User
	// hashing password
	password := input.Password
	email := input.Email
	username := input.Username
	hashedPassword, err := a.hashPassword(password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.Email = email

	// username
	if username == nil {
		username = &strings.Split(email, "@")[0]
		user.Username = *username

	} else {
		user.Username = *input.Username
	}

	if err := a.dbConn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (a *authRepository) ChangePassword(input *dto.ChangePasswordReq, context ...*fiber.Ctx) error {

	newPassword := input.NewPassword
	oldPassword := input.OldPassword
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	userId := a.GetValueFromToken("user_id", c)
	user, err := a.GetUserById(userId)
	util.HandleFunc(err)

	if !a.CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("invalid Credentials")
	}

	hashedNewPassword, err := a.hashPassword(newPassword)
	util.HandleFunc(err)

	user.Password = hashedNewPassword
	a.dbConn.Save(&user)

	return nil
}

func (a *authRepository) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (a *authRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := a.dbConn.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *authRepository) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *authRepository) ValidToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func (a *authRepository) GetUserById(id int) (*entity.User, error) {
	var user entity.User
	if err := a.dbConn.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *authRepository) GetValueFromToken(key string, c *fiber.Ctx) int {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	value := int(claims[key].(float64))
	return value
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{dbConn: db}
}
