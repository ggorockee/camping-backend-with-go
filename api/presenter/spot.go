package presenter

import (
	"camping-backend-with-go/pkg/entities"
)

type Spot struct {
	Id       uint   `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Location string `json:"location"`
}

func SpotSuccessResponse(data *entities.Spot) *JsonResponse {
	spot := Spot{
		Id:       data.Id,
		Title:    data.Title,
		Author:   data.Author,
		Location: data.Location,
	}

	return &JsonResponse{
		Status: true,
		Data:   spot,
		Error:  "",
	}
}

func SpotsSuccessResponse(data *[]Spot) *JsonResponse {
	return &JsonResponse{
		Status: true,
		Data:   data,
		Error:  "",
	}
}

func SpotErrorResponse(err error) *JsonResponse {
	return &JsonResponse{
		Status: false,
		Data:   nil,
		Error:  err.Error(),
	}
}
