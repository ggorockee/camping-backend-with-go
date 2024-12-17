package serializer

import (
	"camping-backend-with-go/internal_backup/domain"
	"camping-backend-with-go/pkg/dto"
)

type UserRole string

const (
	Client UserRole = "client"
	Owner  UserRole = "owner"
	Staff  UserRole = "staff"
	Admin  UserRole = "admin"
)

type UserSerializer interface {
	TinyUserSerialize() dto.TinyUserOut
	UserDetailSerialize() dto.UserDetailOut
}

type userSerializer struct {
	User *entities.User
}

func (u *userSerializer) TinyUserSerialize() dto.TinyUserOut {
	return dto.TinyUserOut{
		Id:       int(u.User.Id),
		Email:    u.User.Email,
		Username: u.User.Username,
		Role:     u.User.Role,
	}
}

func (u *userSerializer) UserDetailSerialize() dto.UserDetailOut {
	return dto.UserDetailOut{
		Id:       int(u.User.Id),
		Email:    u.User.Email,
		Username: u.User.Username,
		Role:     u.User.Role,
	}
}

func NewUserSerializer(u *entities.User) UserSerializer {
	return &userSerializer{User: u}
}
