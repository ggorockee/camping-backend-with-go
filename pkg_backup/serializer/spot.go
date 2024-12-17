package serializer

import (
	entities2 "camping-backend-with-go/internal_backup/domain"
	"camping-backend-with-go/pkg/dto"
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

// Serializer
func SpotsSerializer(spots []entities2.Spot, db *gorm.DB, c *fiber.Ctx) []dto.SpotListOut {
	spotListsRes := make([]dto.SpotListOut, 0)
	for _, spot := range spots {
		userSerializer := NewUserSerializer(&spot.User)
		categorySerializer := NewCategorySerializer(&spot.Category)
		spotSerializer := NewSpotSerializer(&spot, userSerializer, categorySerializer)
		spotListOut := spotSerializer.ListSerialize(db, c)

		if spotListOut.User == nil || spotListOut.User.Id == 0 {
			spotListOut.User = nil
		}

		spotListsRes = append(spotListsRes, spotListOut)
	}
	return spotListsRes
}

type spotSerializer struct {
	spot     *entities2.Spot
	user     UserSerializer
	category CategorySerializer
}

func (s *spotSerializer) GetIsOwner(ctx *fiber.Ctx) bool {
	requestUser, ok := ctx.Locals("request_user").(entities2.User)
	if !ok {
		return false
	}

	return s.spot.User == requestUser
}

func (s *spotSerializer) serializeAmenities() []dto.AmenityListOut {
	amenityListOuts := make([]dto.AmenityListOut, 0)
	for _, amenity := range s.spot.Amenities {
		amenitySerializer := NewAmenitySerializer(&amenity)
		amenityListOuts = append(amenityListOuts, amenitySerializer.ListSerialize())
	}
	return amenityListOuts
}

func (s *spotSerializer) fetchReviews(db *gorm.DB) []entities2.Review {
	var reviews []entities2.Review
	err := db.Where("spot_id = ?", s.spot.Id).Preload("Spot").Find(&reviews).Error
	if err != nil {
		log.Fatalf("DetailSerializer Error, cannot fetch spot instance, %s\n", err.Error())
	}
	return reviews
}

func (s *spotSerializer) calculateAverageRating(reviews []entities2.Review) float64 {
	count := len(reviews)
	if count == 0 {
		return 0
	}
	totalRating := 0.0
	for _, review := range reviews {
		totalRating += float64(review.Rating)
	}
	rating := totalRating / float64(count)
	return math.Round(rating*100) / 100
}

func (s *spotSerializer) makeContext(contexts []*fiber.Ctx) *fiber.Ctx {
	ctx := MakeContext(contexts)
	if ctx == nil {
		log.Fatalf("failed loading fiber.Ctx...")
	}
	return ctx
}

func (s *spotSerializer) ListSerialize(db *gorm.DB, contexts ...*fiber.Ctx) dto.SpotListOut {

	//amenityListOuts := make([]dto.AmenityListOut, 0)
	amenityListOuts := s.serializeAmenities()
	reviews := s.fetchReviews(db)
	rating := s.calculateAverageRating(reviews)
	ctx := s.makeContext(contexts)

	var user *dto.TinyUserOut
	if s.spot.User.Id != 0 {
		tinyUser := s.user.TinyUserSerialize()
		user = &tinyUser
	}

	return dto.SpotListOut{
		Id:          int(s.spot.Id),
		User:        user,
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

	var reviews []entities2.Review
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

func NewSpotSerializer(s *entities2.Spot, u UserSerializer, c CategorySerializer) SpotSerializer {
	return &spotSerializer{spot: s, user: u, category: c}
}
