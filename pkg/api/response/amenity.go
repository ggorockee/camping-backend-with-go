package response

import (
	"camping-backend-with-go/internal/domain/entity"
	"time"
)

type AmenityDetailResponse struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name" gorm:"type:varchar(20);"`
	Description string    `json:"description"`

	// test
	User TinyUserAdapter `json:"user"`
}

type AmenityDetailResponseSlice []AmenityDetailResponse

type AmenityResponseBuilder struct {
	amenities []*entity.Amenity
}

func NewAmenityResponseBuilder() *AmenityResponseBuilder {
	return &AmenityResponseBuilder{}
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
	return AmenityDetailResponse{
		Id:          amenity.Id,
		CreatedAt:   amenity.CreatedAt,
		UpdatedAt:   amenity.UpdatedAt,
		Name:        amenity.Name,
		Description: *amenity.Description,
		User:        TinyUserAdapter{}, // 여기에 사용자 정보를 채워넣어야 합니다
	}
}

func (b *AmenityResponseBuilder) buildSlice() AmenityDetailResponseSlice {
	var result AmenityDetailResponseSlice
	for _, amenity := range b.amenities {
		result = append(result, b.buildSingle(amenity))
	}
	return result
}
