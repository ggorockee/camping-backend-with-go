package entities

import "time"

type Amenity struct {
	Id          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AmenitySerializer interface {
	ListSerialize() AmenityListOut
}

func (a *Amenity) ListSerialize() AmenityListOut {
	return AmenityListOut{
		Id:          int(a.Id),
		Name:        a.Name,
		Description: a.Description,
	}
}

type AmenityListOut struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CreateAmenityInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
