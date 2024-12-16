package serializer

import (
	"camping-backend-with-go/internal/domain"
	"camping-backend-with-go/pkg/dto"
)

type AmenitySerializer interface {
	ListSerialize() dto.AmenityListOut
	DetailSerialize() dto.AmenityDetailOut
}

type amenitySerializer struct {
	Amenity *entities.Amenity
}

func (a *amenitySerializer) ListSerialize() dto.AmenityListOut {
	return dto.AmenityListOut{
		Id:          int(a.Amenity.Id),
		Name:        a.Amenity.Name,
		Description: *a.Amenity.Description,
	}
}

func (a *amenitySerializer) DetailSerialize() dto.AmenityDetailOut {
	return dto.AmenityDetailOut{
		Id:          int(a.Amenity.Id),
		Name:        a.Amenity.Name,
		Description: *a.Amenity.Description,
		CreatedAt:   a.Amenity.CreatedAt,
		UpdatedAt:   a.Amenity.UpdatedAt,
	}
}

func NewAmenitySerializer(a *entities.Amenity) AmenitySerializer {
	return &amenitySerializer{
		Amenity: a,
	}
}
