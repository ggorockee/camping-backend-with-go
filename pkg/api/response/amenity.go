package response

import (
	"camping-backend-with-go/internal/domain/entity"
	"time"
)

type AmenityDetailResponse struct {
	Id          string          `json:"id"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	User        TinyUserReponse `json:"user"`
}

type AmenityDetailResponseSlice []AmenityDetailResponse

type AmenityResponseBuilder struct {
	amenities []*entity.Amenity
	userFunc  func(*entity.Amenity) TinyUserReponse
}

func NewAmenityResponseBuilder() *AmenityResponseBuilder {
	return &AmenityResponseBuilder{
		userFunc: defaultUserFunc,
	}
}

func defaultUserFunc(amenity *entity.Amenity) TinyUserReponse {
	// 기본적으로 빈 사용자 정보를 반환합니다.
	return TinyUserReponse{}
}

func (b *AmenityResponseBuilder) WithUserFunc(f func(*entity.Amenity) TinyUserReponse) *AmenityResponseBuilder {
	b.userFunc = f
	return b
}

func (b *AmenityResponseBuilder) AddAmenity(amenity *entity.Amenity) *AmenityResponseBuilder {
	b.amenities = append(b.amenities, amenity)
	return b
}

func (b *AmenityResponseBuilder) AddAmenities(amenities []*entity.Amenity) *AmenityResponseBuilder {
	b.amenities = append(b.amenities, amenities...)
	return b
}

func (b *AmenityResponseBuilder) Build() interface{} {
	if len(b.amenities) == 1 {
		return b.buildSingle(b.amenities[0])
	}
	return b.buildSlice()
}

func (b *AmenityResponseBuilder) buildSingle(amenity *entity.Amenity) AmenityDetailResponse {
	description := ""
	if amenity.Description != nil {
		description = *amenity.Description
	}

	return AmenityDetailResponse{
		Id:          amenity.Id,
		CreatedAt:   amenity.CreatedAt,
		UpdatedAt:   amenity.UpdatedAt,
		Name:        amenity.Name,
		Description: description,
		User:        b.userFunc(amenity),
	}
}

func (b *AmenityResponseBuilder) buildSlice() AmenityDetailResponseSlice {
	var result AmenityDetailResponseSlice
	for _, amenity := range b.amenities {
		result = append(result, b.buildSingle(amenity))
	}
	return result
}

func GetAmenityResponse(userFunc func(*entity.Amenity) TinyUserReponse, amenities ...*entity.Amenity) interface{} {
	builder := NewAmenityResponseBuilder().WithUserFunc(userFunc)
	return builder.AddAmenities(amenities).Build()
}
