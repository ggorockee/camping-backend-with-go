package user

import (
	"camping-backend-with-go/pkg/config"
	"camping-backend-with-go/pkg/entities"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type Repository interface {
	CreateUser(signUpInputSchema *entities.SignUpInputSchema) error
	hashPassword(password string) (string, error)
	Login(loginInputSchema *entities.LoginInputSchema) (string, error)
	GetUserByEmail(email string) (*entities.User, error)
	CheckPasswordHash(password, hash string) bool
	//validToken(t *jwt.Token, id string) bool
	//validUser(id string, password string) bool
	//GetUserById(id int) (*entities.User, error)
	//CheckPasswordHash(password, hash string) bool
	//getUserByEmail(e string) (*model.User, error)
	//getUserByUsername(u string) (*model.User, error)
}

type repository struct {
	DBConn *gorm.DB
}

func NewRepo(dbconn *gorm.DB) Repository {
	return &repository{
		DBConn: dbconn,
	}
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

func (r *repository) Login(loginInputSchema *entities.LoginInputSchema) (string, error) {
	email := loginInputSchema.Email
	password := loginInputSchema.Password

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

func (r *repository) CreateUser(signUpInputSchema *entities.SignUpInputSchema) error {
	var user entities.User
	// hashing password
	password := signUpInputSchema.Password
	email := signUpInputSchema.Email
	username := signUpInputSchema.Username
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
		user.Username = *signUpInputSchema.Username
	}

	if err := r.DBConn.Create(user).Error; err != nil {
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
