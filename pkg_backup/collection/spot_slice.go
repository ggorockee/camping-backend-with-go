package collection

import (
	"camping-backend-with-go/internal_backup/domain"
	"camping-backend-with-go/pkg/interfaces"
)

type spotSlice []domain.Spot

func (s *spotSlice) Filter(id int) interfaces.SpotCollection {
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

func (s *spotSlice) Remove(spot interfaces.Spot) {
	result := make(spotSlice, 0)
	for _, s := range *s {
		if s.Id != spot.GetId() {
			result = append(result, s)
		}
	}
	*s = result
}

func (s *spotSlice) Add(spot interfaces.Spot) {
	if concreteSpot, ok := spot.(*domain.Spot); ok {
		*s = append(*s, *concreteSpot)
	} else {
		// 타입 단언 실패 시 처리 (예: 로그 출력 또는 패닉)
		panic("invalid Spot type")
	}
}

func (s *spotSlice) ToSlice() []interfaces.Spot {
	result := make([]interfaces.Spot, len(*s))
	for i, spot := range *s {
		result[i] = &spot
	}
	return result
}

func NewSpotCollection(spots []domain.Spot) interfaces.SpotCollection {
	s := spotSlice(spots)
	return &s
}
