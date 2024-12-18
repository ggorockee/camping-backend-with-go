package spotrepository

import (
	reviewdto "camping-backend-with-go/internal/application/dto/review"
	spotdto "camping-backend-with-go/internal/application/dto/spot"
	"camping-backend-with-go/internal/domain/entity"
	amenityrepository "camping-backend-with-go/internal/domain/repository/amenity"
	categoryrepository "camping-backend-with-go/internal/domain/repository/category"
	userrepository "camping-backend-with-go/internal/domain/repository/user"

	"camping-backend-with-go/pkg/util"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type SpotRepository interface {
	CreateSpot(input *spotdto.CreateSpotReq, context ...*fiber.Ctx) (*entity.Spot, error)
	UpdateSpot(input *spotdto.UpdateSpotReq, id int, context ...*fiber.Ctx) (*entity.Spot, error)
	GetSpotById(id int, context ...*fiber.Ctx) (*entity.Spot, error)
	DeleteSpot(id int, context ...*fiber.Ctx) error
	GetAllSpots(context ...*fiber.Ctx) (*[]entity.Spot, error)
	GetReviewsFromSpot(spot *entity.Spot, context ...*fiber.Ctx) (*[]entity.Review, error)
	CreateSpotReview(input *reviewdto.CreateSpotReviewReq, spot *entity.Spot, context ...*fiber.Ctx) (*entity.Review, error)
}

type spotRepository struct {
	dbConn       *gorm.DB
	userRepo     userrepository.UserRepository
	categoryRepo categoryrepository.CategoryRepository
	amenityRepo  amenityrepository.AmenityRepository
}

func (r *spotRepository) CreateSpot(input *spotdto.CreateSpotReq, context ...*fiber.Ctx) (*entity.Spot, error) {
	// jwt에서 user 불러오기
	c, err := util.ContextParser(context...)
	userId := r.userRepo.GetValueFromToken("user_id", c)
	owner, err := r.userRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	// spot
	spot := entity.Spot{
		UserId:      userId, // foreginKey
		User:        *owner, // foreginKey
		Name:        input.Name,
		Country:     input.Country,
		City:        input.City,
		Price:       input.Price,
		Description: *input.Description,
		Address:     input.Address,
		PetFriendly: input.PetFriendly,
		CategoryId:  &input.Category,    // foreginKey
		Category:    entity.Category{},  // foreginKey
		Amenities:   []entity.Amenity{}, // many2many
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// foreignKey
	spot.User = *owner
	spot.UserId = userId

	// category
	categoryObj, err := r.categoryRepo.GetCategoryById(input.Category)
	if err != nil {
		return nil, err
	}

	spot.Category = *categoryObj

	// spot.CategoryId = &input.Category

	// many2many
	amenities := make([]entity.Amenity, 0)
	if input.Amenities != nil {
		for _, amenityId := range *input.Amenities {
			amenity, _ := r.amenityRepo.GetAmenityById(amenityId)
			amenities = append(amenities, *amenity)
		}
	}

	spot.Amenities = amenities

	// value
	if err := r.dbConn.Create(&spot).Error; err != nil {
		return nil, err
	}
	return &spot, nil
}

func (r *spotRepository) UpdateSpot(input *spotdto.UpdateSpotReq, id int, context ...*fiber.Ctx) (*entity.Spot, error) {
	c, err := util.ContextParser(context...)
	userId := r.userRepo.GetValueFromToken("user_id", c)
	spot, err := r.GetSpotById(id)
	if err != nil {
		return nil, err
	}

	fetchedSpotUserId := int(spot.UserId)

	if userId != fetchedSpotUserId {
		return nil, errors.New("permission denied")
	}

	spot.UpdatedAt = time.Now()
	if err := r.dbConn.Model(spot).Updates(input).Error; err != nil {
		return nil, err
	}

	return spot, nil
}

func (r *spotRepository) GetSpotById(id int, context ...*fiber.Ctx) (*entity.Spot, error) {
	var spot entity.Spot

	if err := r.dbConn.Preload("User").Where("id = ?", id).First(&spot).Error; err != nil {
		return nil, err
	}
	return &spot, nil
}

func (r *spotRepository) DeleteSpot(id int, context ...*fiber.Ctx) error {
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	userId := r.userRepo.GetValueFromToken("user_id", c)
	spot, err := r.GetSpotById(id)
	if err != nil {
		return err
	}

	if int(spot.UserId) != userId {
		return errors.New("permission denied")
	}

	if err := r.dbConn.Delete(spot).Error; err != nil {
		return err
	}

	return nil
}

func (r *spotRepository) GetAllSpots(context ...*fiber.Ctx) (*[]entity.Spot, error) {
	var spots []entity.Spot
	if err := r.dbConn.Preload("User").Find(&spots).Error; err != nil {
		return nil, err
	}

	return &spots, nil
}

func (r *spotRepository) GetReviewsFromSpot(spot *entity.Spot, context ...*fiber.Ctx) (*[]entity.Review, error) {
	var reviews []entity.Review

	if err := r.dbConn.Where("spot_id = ?", spot.Id).Preload("Spot").Find(&reviews).Error; err != nil {
		return nil, err
	}

	return &reviews, nil
}

func (r *spotRepository) CreateSpotReview(input *reviewdto.CreateSpotReviewReq, spot *entity.Spot, context ...*fiber.Ctx) (*entity.Review, error) {
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	requestUser, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return nil, errors.New("use is not authenticated")
	}

	// spot은 이미 불러와짐
	review := entity.Review{
		Id:        0,
		User:      requestUser,
		Spot:      *spot,
		Payload:   input.Payload,
		Rating:    input.Rating,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := r.dbConn.Create(&review).Error; err != nil {
		return nil, err
	}

	return &review, nil
}

func NewSpotRepository(
	db *gorm.DB,
	u userrepository.UserRepository,
	c categoryrepository.CategoryRepository,
	a amenityrepository.AmenityRepository,
) SpotRepository {
	return &spotRepository{
		dbConn:       db,
		userRepo:     u,
		categoryRepo: c,
		amenityRepo:  a,
	}
}
