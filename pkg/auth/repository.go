package auth

import (
	"camping-backend-with-go/pkg/entities"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	Login(loginInput *entities.Login) (string, error)
	CheckPasswordHash(password, hash string) bool
	GetUserByEmail(email string) (*entities.User, error)
}

type repository struct {
	DBConn *gorm.DB
}

func NewRepo(dbconn *gorm.DB) Repository {
	return &repository{DBConn: dbconn}
}

func (r *repository) Login(loginInput *entities.Login) (string, error) {
	user, err := r.GetUserByEmail(loginInput.Email)
	if err != nil {
		return "", err
	}

	if !r.CheckPasswordHash(loginInput.Password, user.Password) {
		return "", errors.New("password didn't not match")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id

	t, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}

	return t, nil
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
