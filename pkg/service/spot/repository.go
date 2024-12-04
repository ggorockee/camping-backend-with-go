package spot

import (
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/user"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository interface {
	CreateSpot(spot *entities.CreateSpotInputSchema, ctx *fiber.Ctx) (*entities.Spot, error)
	FetchMySpots(ctx *fiber.Ctx) (*[]entities.Spot, error)
	GetSpot(id int, ctx *fiber.Ctx) (*entities.Spot, error)
	UpdateSpot(spot *entities.Spot, id int) (*entities.Spot, error)
	GetFindById(id int) (*entities.Spot, error)
	PartialUpdateSpot(spot *entities.Spot, id int) (*entities.Spot, error)
	DeleteSpot(id int) error
	GetSpotById(id int) (*entities.Spot, error)
}

type repository struct {
	DBConn   *gorm.DB
	UserRepo user.Repository
}

func NewRepo(dbconn *gorm.DB, userRepo user.Repository) Repository {
	return &repository{
		DBConn:   dbconn,
		UserRepo: userRepo,
	}
}

// GetSpotById implements Repository.
func (r *repository) GetSpotById(id int) (*entities.Spot, error) {
	var spot entities.Spot

	if err := r.DBConn.Preload("User").Where("id = ?", id).Find(&spot).Error; err != nil {
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

func (r *repository) CreateSpot(createSpotInputSchema *entities.CreateSpotInputSchema, ctx *fiber.Ctx) (*entities.Spot, error) {
	title := createSpotInputSchema.Title
	location := createSpotInputSchema.Location

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

func (r *repository) UpdateSpot(spot *entities.Spot, id int) (*entities.Spot, error) {
	fetched, err := r.GetFindById(id)

	if err != nil {
		return nil, err
	}

	spot.UpdatedAt = time.Now()

	result := r.DBConn.Model(fetched).Updates(&spot)

	if result.Error != nil {
		return nil, result.Error
	}

	return spot, nil
}

func (r *repository) PartialUpdateSpot(spot *entities.Spot, id int) (*entities.Spot, error) {
	fetched, err := r.GetFindById(id)
	if err != nil {
		return nil, err
	}

	spot.UpdatedAt = time.Now()

	result := r.DBConn.Model(fetched).Updates(&spot)
	if result.Error != nil {
		return nil, result.Error
	}

	return fetched, nil
}

func (r *repository) DeleteSpot(id int) error {
	spot, err := r.GetFindById(id)
	if err != nil {
		return err
	}

	result := r.DBConn.Delete(spot)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
