package serializer

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"time"
)

type SpotSerializer interface {
	ListSerialize() dto.SpotListOut
	DetailSerialize() dto.SpotDetailOut
}

type spotSerializer struct {
	spot     *entities.Spot
	user     UserSerializer
	category CategorySerializer
}

func (s *spotSerializer) ListSerialize() dto.SpotListOut {

	amenityListOuts := make([]dto.AmenityListOut, 0)
	for _, amenity := range s.spot.Amenities {
		amenitySerializer := NewAmenitySerializer(&amenity)
		amenityListOuts = append(amenityListOuts, amenitySerializer.ListSerialize())
	}
	return dto.SpotListOut{
		Id:          int(s.spot.Id),
		User:        s.user.TinyUserSerialize(),
		Name:        s.spot.Name,
		Country:     s.spot.Country,
		City:        s.spot.City,
		Price:       s.spot.Price,
		Description: &s.spot.Description,
		Address:     s.spot.Address,
		PetFriendly: s.spot.PetFriendly,
		Category:    s.category.ListSerialize(),
		Amenities:   &amenityListOuts,

		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func (s *spotSerializer) DetailSerialize() dto.SpotDetailOut {
	amenityListOuts := make([]dto.AmenityListOut, 0)
	for _, amenity := range s.spot.Amenities {
		amenitySerializer := NewAmenitySerializer(&amenity)
		amenityListOuts = append(amenityListOuts, amenitySerializer.ListSerialize())
	}

	return dto.SpotDetailOut{
		Id:          int(s.spot.Id),
		User:        s.user.TinyUserSerialize(),
		Name:        s.spot.Name,
		Country:     s.spot.Country,
		City:        s.spot.City,
		Price:       s.spot.Price,
		Description: s.spot.Description,
		Address:     s.spot.Address,
		PetFriendly: s.spot.PetFriendly,
		Category:    s.category.ListSerialize(),
		Amenities:   &amenityListOuts,

		CreatedAt: s.spot.CreatedAt,
		UpdatedAt: s.spot.UpdatedAt,
	}
}

func NewSpotSerializer(s *entities.Spot, u UserSerializer, c CategorySerializer) SpotSerializer {
	return &spotSerializer{spot: s, user: u, category: c}
}
