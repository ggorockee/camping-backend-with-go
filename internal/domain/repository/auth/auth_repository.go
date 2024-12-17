package authrepository

import (
	authdto "camping-backend-with-go/internal/application/dto/auth"
	userentity "camping-backend-with-go/internal/domain/entity/user"
	"camping-backend-with-go/pkg/util"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
)

type AuthRepository interface {
	Login(input *authdto.LoginReq) (string, error)
	CreateUser(input *authdto.SignUpReq) error
	ChangePassword(input *authdto.ChangePasswordReq, contexts ...*fiber.Ctx) error

	hashPassword(password string) (string, error)
	GetUserByEmail(email string) (*userentity.User, error)
	CheckPasswordHash(password, hash string) bool
	ValidToken(t *jwt.Token, id string) bool
	GetUserById(id int) (*userentity.User, error)
	GetValueFromToken(key string, c *fiber.Ctx) int
}

type authRepository struct {
	dbConn *gorm.DB
}

func (a *authRepository) Login(input *authdto.LoginReq) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authRepository) CreateUser(input *authdto.SignUpReq) error {
	//TODO implement me
	panic("implement me")
}

func (a *authRepository) ChangePassword(input *authdto.ChangePasswordReq, contexts ...*fiber.Ctx) error {

	newPassword := input.NewPassword
	oldPassword := input.OldPassword
	c, err := util.ContextParser(contexts...)
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

func (a *authRepository) GetUserByEmail(email string) (*userentity.User, error) {
	var user userentity.User
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

func (a *authRepository) GetUserById(id int) (*userentity.User, error) {
	var user userentity.User
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
