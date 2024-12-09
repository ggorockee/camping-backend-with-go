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
	spot    *entities.Spot
	user    UserSerializer
	amenity AmenitySerializer
}

func (s *spotSerializer) ListSerialize() dto.SpotListOut {
	amenityIds := make([]int, 0)
	for _, amenity := range s.spot.Amenities {
		amenityIds = append(amenityIds, int(amenity.Id))
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
		Category:    *s.spot.CategoryId,
		Amenities:   amenityIds,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func (s *spotSerializer) DetailSerialize() dto.SpotDetailOut {
	return dto.SpotDetailOut{
		Id:        int(s.spot.Id),
		User:      s.user.TinyUserSerialize(),
		Title:     s.spot.Title,
		Location:  s.spot.Location,
		Author:    s.spot.Author,
		CreatedAt: s.spot.CreatedAt,
		UpdatedAt: s.spot.UpdatedAt,
		Review:    s.spot.Review,
	}
}

func NewSpotSerializer(s *entities.Spot, u UserSerializer, a AmenitySerializer) SpotSerializer {
	return &spotSerializer{spot: s, user: u, amenity: a}
}
