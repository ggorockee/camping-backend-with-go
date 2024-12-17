package user

import (
	"camping-backend-with-go/internal_backup/domain"
	"camping-backend-with-go/pkg/config"
	"camping-backend-with-go/pkg/dto"
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(input *dto.SignUpIn) error
	hashPassword(password string) (string, error)
	Login(input *dto.LoginIn) (string, error)
	GetUserByEmail(email string) (*entities.User, error)
	CheckPasswordHash(password, hash string) bool
	ChangePassword(input *dto.ChangePasswordIn, ctx *fiber.Ctx) error
	ValidToken(t *jwt.Token, id string) bool
	GetUserById(id int) (*entities.User, error)
	GetValueFromToken(key string, ctx *fiber.Ctx) int
	ValidUser(id int, user *entities.User) error
	//validUser(id string, password string) bool
	//CheckPasswordHash(password, hash string) bool
	//getUserByEmail(e string) (*model.User, error)
	//getUserByUsername(u string) (*model.User, error)
}

type repository struct {
	DBConn *gorm.DB
}

func NewRepo(dbConn *gorm.DB) Repository {
	return &repository{
		DBConn: dbConn,
	}
}

func (r *repository) ValidToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func (r *repository) ValidUser(id int, user *entities.User) error {
	if id != int(user.Id) {
		return errors.New("ValidUser:: => invalid user")
	}
	return nil
}

func (r *repository) GetUserById(id int) (*entities.User, error) {
	var user entities.User
	if err := r.DBConn.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetValueFromToken(key string, ctx *fiber.Ctx) int {
	token := ctx.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	value := int(claims[key].(float64))
	return value
}

func (r *repository) ChangePassword(input *dto.ChangePasswordIn, ctx *fiber.Ctx) error {
	newPassword := input.NewPassword
	oldPassword := input.OldPassword

	// GetFindByEmail
	// login한 User의 Id를 알아내는 로직
	// user_id

	userId := r.GetValueFromToken("user_id", ctx)

	user, err := r.GetUserById(userId)

	if err != nil {
		return err
	}

	// CheckPassword
	if !r.CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("invalid Credentials")
	}

	hashedNewPassword, err := r.hashPassword(newPassword)
	if err != nil {
		return err
	}

	// updatePassword
	user.Password = hashedNewPassword

	// Database save
	r.DBConn.Save(&user)
	return nil
}

func (r *repository) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (r *repository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	if err := r.DBConn.Where(entities.User{Email: email}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) Login(input *dto.LoginIn) (string, error) {
	email := input.Email
	password := input.Password

	user, err := r.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !r.CheckPasswordHash(password, user.Password) {
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

func (r *repository) CreateUser(input *dto.SignUpIn) error {
	var user entities.User
	// hashing password
	password := input.Password
	email := input.Email
	username := input.Username
	hashedPassword, err := r.hashPassword(password)
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

	if err := r.DBConn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//func (r *repository) validToken(t *jwt.Token, id string) bool {
//	n, err := strconv.Atoi(id)
//	if err != nil {
//		return false
//	}
//
//	claims := t.Claims.(jwt.MapClaims)
//	uid := int(claims["user_id"].(float64))
//
//	return uid == n
//}
//
//func (r *repository) validUser(id string, password string) bool {
//	var user entities.User
//
//	r.DBConn.First(&user, id)
//	if user.Email == "" {
//		return false
//	}
//
//	if
//}
//
//func (r *repository) GetUserById(id int) (*entities.User, error) {
//	//TODO implement me
//	panic("implement me")
//}
