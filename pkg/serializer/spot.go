package serializer

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"math"
)

type SpotSerializer interface {
	ListSerialize(db *gorm.DB, contexts ...*fiber.Ctx) dto.SpotListOut
	DetailSerialize(db *gorm.DB, contexts ...*fiber.Ctx) dto.SpotDetailOut
	GetIsOwner(ctx *fiber.Ctx) bool
}

type spotSerializer struct {
	spot     *entities.Spot
	user     UserSerializer
	category CategorySerializer
}

func (s *spotSerializer) GetIsOwner(ctx *fiber.Ctx) bool {
	requestUser, ok := ctx.Locals("request_user").(entities.User)
	if !ok {
		return false
	}

	return s.spot.User == requestUser
}

func (s *spotSerializer) ListSerialize(db *gorm.DB, contexts ...*fiber.Ctx) dto.SpotListOut {

	amenityListOuts := make([]dto.AmenityListOut, 0)
	for _, amenity := range s.spot.Amenities {
		amenitySerializer := NewAmenitySerializer(&amenity)
		amenityListOuts = append(amenityListOuts, amenitySerializer.ListSerialize())
	}

	var reviews []entities.Review
	err := db.Where("spot_id = ?", s.spot.Id).Preload("Spot").Find(&reviews).Error
	if err != nil {
		log.Fatalf("DeatilSerializer Error, cannot fetched spot instance, %s\n", err.Error())
	}

	count := len(reviews)
	var rating float64
	if count == 0 {
		rating = 0
	} else {
		totalRating := 0.0
		for _, review := range reviews {
			totalRating += float64(review.Rating)
		}
		rating = totalRating / float64(count)
		rating = math.Round(rating*100) / 100
	}

	ctx := MakeContext(contexts)
	if ctx == nil {
		log.Fatalf("failed loading fiber.Ctx...")
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
		Rating:      rating,
		IsOwner:     s.GetIsOwner(ctx),

		CreatedAt: s.spot.CreatedAt,
		UpdatedAt: s.spot.UpdatedAt,
	}
}

func (s *spotSerializer) DetailSerialize(db *gorm.DB, contexts ...*fiber.Ctx) dto.SpotDetailOut {
	amenityListOuts := make([]dto.AmenityListOut, 0)
	for _, amenity := range s.spot.Amenities {
		amenitySerializer := NewAmenitySerializer(&amenity)
		amenityListOuts = append(amenityListOuts, amenitySerializer.ListSerialize())
	}

	var reviews []entities.Review
	err := db.Where("spot_id = ?", s.spot.Id).Preload("Spot").Find(&reviews).Error
	if err != nil {
		log.Fatalf("DeatilSerializer Error, cannot fetched spot instance, %s\n", err.Error())
	}

	count := len(reviews)
	var rating float64
	if count == 0 {
		rating = 0
	} else {
		totalRating := 0.0
		for _, review := range reviews {
			totalRating += float64(review.Rating)
		}
		rating = totalRating / float64(count)
		rating = math.Round(rating*100) / 100
	}

	// context
	ctx := MakeContext(contexts)
	if ctx == nil {
		log.Fatalf("failed loading fiber.Ctx...")
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
		Rating:      rating,
		IsOwner:     s.GetIsOwner(ctx),

		CreatedAt: s.spot.CreatedAt,
		UpdatedAt: s.spot.UpdatedAt,
	}
}

func NewSpotSerializer(s *entities.Spot, u UserSerializer, c CategorySerializer) SpotSerializer {
	return &spotSerializer{spot: s, user: u, category: c}
}
