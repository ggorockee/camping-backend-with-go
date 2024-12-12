package serializer

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"time"
)

type ReviewSerializer interface {
	Serialize() ReviewOut
}

type ReviewsSerializer interface {
	Serialize() []ReviewOut
}

type ReviewOut struct {
	Id uint `json:"id" gorm:"primaryKey"`
	// UserId int  `json:"user_id"`
	//User   entities.User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	User dto.TinyUserOut

	// SpotId *int `json:"spot_id"`
	//Spot   entities.Spot `gorm:"foreignKey:SpotId;constraint:OnDelete:SET NULL"`

	Payload string `json:"payload"`
	Rating  int    `json:"rating"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type reviewSerializer struct {
	review *entities.Review
	user   UserSerializer
}

type reviewsSerializer struct {
	reviews *[]entities.Review
}

func (r *reviewsSerializer) Serialize() []ReviewOut {
	// 다중 객체 serializer
	var serializedReviews []ReviewOut
	for _, review := range *r.reviews {
		serializedUser := NewUserSerializer(&review.User)
		serializedReview := NewReviewSerializer(&review, serializedUser)
		serializedReviews = append(serializedReviews, serializedReview.Serialize())
	}

	return serializedReviews
}

func (r *reviewSerializer) Serialize() ReviewOut {
	// 단일 객체 serializer
	return ReviewOut{
		Id: r.review.Id,
		// UserId: r.review.UserId,
		User: r.user.TinyUserSerialize(),
		// SpotId: r.review.SpotId,
		//Spot:      r.review.Spot,
		Payload:   r.review.Payload,
		Rating:    r.review.Rating,
		CreatedAt: r.review.CreatedAt,
		UpdatedAt: r.review.UpdatedAt,
	}
}

func NewReviewSerializer(review *entities.Review, user UserSerializer) ReviewSerializer {
	return &reviewSerializer{review: review, user: user}
}

func NewReviewsSerializer(reviews *[]entities.Review) ReviewsSerializer {

	return &reviewsSerializer{reviews: reviews}
}
