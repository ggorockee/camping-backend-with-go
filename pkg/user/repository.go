package user

import (
	"camping-backend-with-go/pkg/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	hashPassword(password string) (string, error)
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

func (r *repository) CreateUser(user *entities.User) (*entities.User, error) {
	hashedPassword, err := r.hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	result := r.DBConn.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
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
