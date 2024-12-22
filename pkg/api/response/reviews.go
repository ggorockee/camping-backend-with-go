package response

import (
	"camping-backend-with-go/internal/domain/entity"
	"encoding/json"
	"time"
)

// type ReviewTinyResponse struct {

// }

type ReviewDetailResponse struct {
	Id        string          `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	User      TinyUserAdapter `json:"user"`

	Payload string `json:"payload"`
	Rating  int    `json:"rating"`
}

type ReviewDetailAdapter struct {
	*ReviewDetailResponse
}

func NewReviewDetailAdapter(reviewEntity *entity.Review, userEntity *entity.User) *ReviewDetailAdapter {

	userAdapter := NewTinyUserAdapter(userEntity)

	reviewDetail := &ReviewDetailResponse{
		Id:        reviewEntity.Id,
		CreatedAt: reviewEntity.CreatedAt,
		UpdatedAt: reviewEntity.UpdatedAt,
		User:      *userAdapter,
		Payload:   reviewEntity.Payload,
		Rating:    reviewEntity.Rating,
	}
	return &ReviewDetailAdapter{ReviewDetailResponse: reviewDetail}
}

func (a *ReviewDetailAdapter) MarshalJSON() ([]byte, error) {
	userResponse := ReviewDetailResponse{
		Id:        a.Id,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		User:      a.User,
		Payload:   a.Payload,
		Rating:    a.Rating,
	}
	return json.Marshal(userResponse)
}
