package spot

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/amenity"
	"camping-backend-with-go/pkg/service/category"
	"camping-backend-with-go/pkg/service/user"
	"camping-backend-with-go/pkg/service/util"
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository interface {
	CreateSpot(input *dto.CreateSpotIn, ctx *fiber.Ctx) (*entities.Spot, error)
	FetchMySpots(ctx *fiber.Ctx) (*[]entities.Spot, error)
	GetSpot(id int, ctx ...*fiber.Ctx) (*entities.Spot, error)
	UpdateSpot(input *dto.UpdateSpotIn, id int, ctx *fiber.Ctx) (*entities.Spot, error)
	GetFindById(id int) (*entities.Spot, error)
	DeleteSpot(id int, ctx *fiber.Ctx) error
	GetSpotById(id int) (*entities.Spot, error)
	GetAllSpots() (*[]entities.Spot, error)
	GetReviewsFromSpot(spot *entities.Spot, contexts ...*fiber.Ctx) (*[]entities.Review, error)
	CreateSpotReview(input *dto.CreateSpotReviewReq, spot *entities.Spot, contexts ...*fiber.Ctx) (*entities.Review, error)
}

type repository struct {
	DBConn      *gorm.DB
	UserRepo    user.Repository
	AmenityRepo amenity.Repository
	CatRepo     category.Repository
}

// CreateSpotReview implements Repository.
func (r *repository) CreateSpotReview(input *dto.CreateSpotReviewReq, spot *entities.Spot, contexts ...*fiber.Ctx) (*entities.Review, error) {
	c, err := util.ContextParser(contexts...)
	if err != nil {
		return nil, err
	}

	requestUser, ok := c.Locals("request_user").(entities.User)
	if !ok {
		return nil, errors.New("use is not authenticated")
	}

	// spot은 이미 불러와짐
	review := entities.Review{
		Id:        0,
		User:      requestUser,
		Spot:      *spot,
		Payload:   input.Payload,
		Rating:    input.Rating,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := r.DBConn.Create(&review).Error; err != nil {
		return nil, err
	}

	return &review, nil

}

func (r *repository) GetReviewsFromSpot(spot *entities.Spot, contexts ...*fiber.Ctx) (*[]entities.Review, error) {
	var reviews []entities.Review

	if err := r.DBConn.Where("spot_id = ?", spot.Id).Preload("Spot").Find(&reviews).Error; err != nil {
		return nil, err
	}

	return &reviews, nil
}

func NewRepo(
	dbConn *gorm.DB,
	userRepo user.Repository,
	amenRepo amenity.Repository,
	catRepo category.Repository,
) Repository {
	return &repository{
		DBConn:      dbConn,
		UserRepo:    userRepo,
		AmenityRepo: amenRepo,
		CatRepo:     catRepo,
	}
}

// GetAllSpots implements Repository.
func (r *repository) GetAllSpots() (*[]entities.Spot, error) {
	var spots []entities.Spot
	if err := r.DBConn.Preload("User").Find(&spots).Error; err != nil {
		return nil, err
	}

	return &spots, nil
}

// GetSpotById implements Repository.
func (r *repository) GetSpotById(id int) (*entities.Spot, error) {
	var spot entities.Spot

	if err := r.DBConn.Preload("User").Where("id = ?", id).First(&spot).Error; err != nil {
		return nil, err
	}

	return &spot, nil
}

func (r *repository) GetSpot(id int, ctx ...*fiber.Ctx) (*entities.Spot, error) {

	fetchedSpot, err := r.GetSpotById(id)
	if err != nil {
		return nil, err
	}
	return fetchedSpot, nil
}

func (r *repository) GetFindById(id int) (*entities.Spot, error) {
	var spot entities.Spot
	result := r.DBConn.First(&spot, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &spot, nil
}

func (r *repository) CreateSpot(input *dto.CreateSpotIn, ctx *fiber.Ctx) (*entities.Spot, error) {
	// jwt에서 user 불러오기
	userId := r.UserRepo.GetValueFromToken("user_id", ctx)
	owner, err := r.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	// spot
	spot := entities.Spot{
		UserId:      userId, // foreginKey
		User:        *owner, // foreginKey
		Name:        input.Name,
		Country:     input.Country,
		City:        input.City,
		Price:       input.Price,
		Description: *input.Description,
		Address:     input.Address,
		PetFriendly: input.PetFriendly,
		CategoryId:  &input.Category,      // foreginKey
		Category:    entities.Category{},  // foreginKey
		Amenities:   []entities.Amenity{}, // many2many
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// foreignKey
	spot.User = *owner
	spot.UserId = userId

	// category
	categoryObj, err := r.CatRepo.GetCategoryById(input.Category)
	if err != nil {
		return nil, err
	}

	spot.Category = *categoryObj

	// spot.CategoryId = &input.Category

	// many2many
	amenities := make([]entities.Amenity, 0)
	if input.Amenities != nil {
		for _, amenityId := range *input.Amenities {
			amentyObj, _ := r.AmenityRepo.GetAmenityById(amenityId)
			amenities = append(amenities, *amentyObj)
		}
	}

	spot.Amenities = amenities

	// value
	if err := r.DBConn.Create(&spot).Error; err != nil {
		return nil, err
	}

	return &spot, nil
}

func (r *repository) FetchMySpots(ctx *fiber.Ctx) (*[]entities.Spot, error) {
	// Login이 되어있어야함
	// 0. middleware 처리 (v)
	// 1. jwtToken을 가지고와서 userId를 얻음(from localstorage)
	userId := r.UserRepo.GetValueFromToken("user_id", ctx)

	// 2. userId를 이용한 query
	var spots []entities.Spot
	if err := r.DBConn.Preload("User").Where("user_id = ?", userId).Find(&spots).Error; err != nil {
		return nil, err
	}

	return &spots, nil
}

func (r *repository) UpdateSpot(input *dto.UpdateSpotIn, id int, ctx *fiber.Ctx) (*entities.Spot, error) {
	// Login이 되어있어야함
	// 0. middleware 처리 (v)
	// 1. jwtToken을 가지고와서 userId를 얻음(from localstorage)
	// 2. :id로 불러온 spot.user_id와 jwtToken값이 같아야함
	// 3. update

	userId := r.UserRepo.GetValueFromToken("user_id", ctx)
	fetchedSpot, err := r.GetSpotById(id)
	if err != nil {
		return nil, err
	}

	fetchedSpotUserId := int(fetchedSpot.UserId)

	if userId != fetchedSpotUserId {
		return nil, errors.New("permission denied")
	}

	fetchedSpot.UpdatedAt = time.Now()
	if err := r.DBConn.Model(fetchedSpot).Updates(input).Error; err != nil {
		return nil, err
	}

	return fetchedSpot, nil

}

func (r *repository) DeleteSpot(id int, ctx *fiber.Ctx) error {
	// Login이 되어있어야함
	// 0. middleware 처리 (v)
	userId := r.UserRepo.GetValueFromToken("user_id", ctx)
	// 1. jwtToken을 가지고와서 userId를 얻음(from localstorage)

	// 2. :id로 불러온 spot.user_id와 jwtToken값이 같아야함
	fetchedSpot, err := r.GetSpotById(id)
	log.Println("fetchedSpot: ", fetchedSpot)
	if err != nil {
		return err
	}

	if int(fetchedSpot.UserId) != userId {
		return errors.New("permission denied")
	}

	if err := r.DBConn.Delete(fetchedSpot).Error; err != nil {
		return err
	}

	return nil
}
