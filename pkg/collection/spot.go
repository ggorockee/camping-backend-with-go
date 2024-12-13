package collection

import "camping-backend-with-go/pkg/entities"

type spotSlice []entities.Spot

func (s *spotSlice) Filter(id int) SpotCollection {
	result := make(spotSlice, 0)
	for _, spot := range *s {
		if int(spot.Id) == id {
			result = append(result, spot)
		}
	}

	return &result
}

func (s *spotSlice) Exists() bool {
	return len(*s) > 0
}

func (s *spotSlice) Remove(spot *entities.Spot) {
	result := make(spotSlice, 0)
	for _, s := range *s {
		if s.Id != spot.Id {
			result = append(result, s)
		}
	}
	*s = result
}

func (s *spotSlice) Add(spot *entities.Spot) {
	*s = append(*s, *spot)
}

func (s *spotSlice) ToSlice() []entities.Spot {
	return *s
}

func NewSpotCollection(spots []entities.Spot) SpotCollection {
	s := spotSlice(spots)
	return &s
}
