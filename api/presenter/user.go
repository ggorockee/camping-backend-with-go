package presenter

import (
	"camping-backend-with-go/pkg/entities"
)

type User struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func UserErrorResponse(err error) *JsonResponse {

	return &JsonResponse{
		Status: false,
		Data:   nil,
		Error:  err.Error(),
	}
}

func UserSuccessResponse(data *entities.User) *JsonResponse {
	user := User{
		Id:       data.Id,
		Email:    data.Email,
		Username: data.Username,
	}
	return &JsonResponse{
		Status: true,
		Data:   user,
		Error:  "",
	}
}
