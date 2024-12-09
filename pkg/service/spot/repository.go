package spot

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/user"
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository interface {
	CreateSpot(input *dto.CreateSpotIn, ctx *fiber.Ctx) (*entities.Spot, error)
	FetchMySpots(ctx *fiber.Ctx) (*[]entities.Spot, error)
	GetSpot(id int, ctx *fiber.Ctx) (*entities.Spot, error)
	UpdateSpot(input *dto.UpdateSpotIn, id int, ctx *fiber.Ctx) (*entities.Spot, error)
	GetFindById(id int) (*entities.Spot, error)
	DeleteSpot(id int, ctx *fiber.Ctx) error
	GetSpotById(id int) (*entities.Spot, error)
	GetAllSpots() (*[]entities.Spot, error)
}

type repository struct {
	DBConn   *gorm.DB
	UserRepo user.Repository
}

func NewRepo(dbConn *gorm.DB, userRepo user.Repository) Repository {
	return &repository{
		DBConn:   dbConn,
		UserRepo: userRepo,
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

// Spot
func (r *repository) GetSpot(id int, ctx *fiber.Ctx) (*entities.Spot, error) {
	// Login이 되어있어야함
	// 0. middleware 처리 (v)
	// 1. jwtToken을 가지고와서 userId를 얻음(from localstorage)
	userId := r.UserRepo.GetValueFromToken("user_id", ctx)

	// 2. userId를 이용해 user instance를 가지고옴
	fetchedUser, err := r.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	// 3. :id를 이용해 spot인스턴스를 가져옴
	fetchedSpot, err := r.GetSpotById(id)
	if err != nil {
		return nil, err
	}
	spotUserId := int(fetchedSpot.UserId)

	// validation
	// 4. spot instance의 userid와 user instance의 id가 같은지 비교
	err = r.UserRepo.ValidUser(spotUserId, fetchedUser)
	if err != nil {
		return nil, err
	}

	// success
	// 5. 모든 과정이 통과되었다면 spot객체 return
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
	title := input.Title
	location := input.Location

	// jwt에서 user 불러오기
	userId := r.UserRepo.GetValueFromToken("user_id", ctx)
	user, err := r.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	// spot
	var spot entities.Spot
	spot.Location = location

	// fill spot value
	spot.Title = title
	spot.UpdatedAt = time.Now()
	spot.CreatedAt = time.Now()
	spot.Author = user.Username
	spot.UserId = uint(userId)
	spot.User = *user
	// 추가되는 것들
	spot.Review = input.Review

	result := r.DBConn.Create(&spot)
	if result.Error != nil {
		return nil, result.Error
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
	fetchedSpotUserId := int(fetchedSpot.UserId)
	if err != nil {
		return nil, err
	}

	if userId != fetchedSpotUserId {
		return nil, errors.New("permission denied")
	}

	updated_title := input.Title
	if updated_title != "" {
		fetchedSpot.Title = updated_title
	}

	updated_location := input.Location
	if updated_location != "" {
		fetchedSpot.Location = updated_location
	}

	if input.Review != "" {
		fetchedSpot.Review = input.Review
	}

	fetchedSpot.UpdatedAt = time.Now()

	if err := r.DBConn.Model(fetchedSpot).Updates(fetchedSpot).Error; err != nil {
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
