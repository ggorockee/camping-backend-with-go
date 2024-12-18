package userrepository

import (
	userdto "camping-backend-with-go/internal/application/dto/user"
	"camping-backend-with-go/internal/domain/entity"
	"camping-backend-with-go/pkg/util"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
)

type UserRepository interface {
	HashPassword(password string, context ...*fiber.Ctx) (string, error)
	GetUserByEmail(email string, context ...*fiber.Ctx) (*entity.User, error)
	CheckPasswordHash(password, hash string, context ...*fiber.Ctx) bool          // auth
	ChangePassword(input *userdto.ChangePasswordReq, context ...*fiber.Ctx) error //
	ValidToken(t *jwt.Token, id string, context ...*fiber.Ctx) bool
	GetUserById(id int, context ...*fiber.Ctx) (*entity.User, error)
	GetValueFromToken(key string, context ...*fiber.Ctx) int
	ValidUser(id int, user *entity.User, context ...*fiber.Ctx) error
}

type userRepository struct {
	dbConn *gorm.DB
}

func (r *userRepository) HashPassword(password string, context ...*fiber.Ctx) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (r *userRepository) GetUserByEmail(email string, context ...*fiber.Ctx) (*entity.User, error) {
	var user entity.User

	if err := r.dbConn.Where(entity.User{Email: email}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CheckPasswordHash(password, hash string, context ...*fiber.Ctx) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (r *userRepository) ChangePassword(input *userdto.ChangePasswordReq, context ...*fiber.Ctx) error {
	c, err := util.ContextParser(context...)
	newPassword := input.NewPassword
	oldPassword := input.OldPassword

	userId := r.GetValueFromToken("user_id", c)
	user, err := r.GetUserById(userId)

	if err != nil {
		return err
	}

	// CheckPassword
	if !r.CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("invalid Credentials")
	}

	hashedNewPassword, err := r.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// updatePassword
	user.Password = hashedNewPassword

	// Database save
	r.dbConn.Save(&user)
	return nil
}

func (r *userRepository) ValidToken(t *jwt.Token, id string, context ...*fiber.Ctx) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func (r *userRepository) GetUserById(id int, context ...*fiber.Ctx) (*entity.User, error) {
	var user entity.User
	if err := r.dbConn.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetValueFromToken(key string, context ...*fiber.Ctx) int {
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	value := int(claims[key].(float64))
	return value
}

func (r *userRepository) ValidUser(id int, user *entity.User, context ...*fiber.Ctx) error {
	if id != int(user.Id) {
		return errors.New("ValidUser:: => invalid user")
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{dbConn: db}
}
