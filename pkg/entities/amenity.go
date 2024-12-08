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
	DetailSerialize() AmenityDetailOut
}

type amenitySerializer struct {
	Amenity *Amenity
}

func (a *amenitySerializer) ListSerialize() AmenityListOut {
	return AmenityListOut{
		Id:          int(a.Amenity.Id),
		Name:        a.Amenity.Name,
		Description: a.Amenity.Description,
	}
}

func (a *amenitySerializer) DetailSerialize() AmenityDetailOut {
	return AmenityDetailOut{
		Id:          int(a.Amenity.Id),
		Name:        a.Amenity.Name,
		Description: a.Amenity.Description,
		CreatedAt:   a.Amenity.CreatedAt,
		UpdatedAt:   a.Amenity.UpdatedAt,
	}
}

func NewAmenitySerializer(a *Amenity) AmenitySerializer {
	return &amenitySerializer{
		Amenity: a,
	}
}

type AmenityListOut struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type AmenityDetailOut struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateAmenityInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type UpdateAmenityInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
