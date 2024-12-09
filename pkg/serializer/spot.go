package serializer

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
)

type SpotSerializer interface {
	ListSerialize() dto.SpotListOut
	DetailSerialize() dto.SpotDetailOut
}

type spotSerializer struct {
	spot *entities.Spot
	user UserSerializer
}

func (s *spotSerializer) ListSerialize() dto.SpotListOut {

	return dto.SpotListOut{
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

func NewSpotSerializer(s *entities.Spot, u UserSerializer) SpotSerializer {
	return &spotSerializer{spot: s, user: u}
}
