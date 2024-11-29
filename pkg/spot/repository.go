package spot

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Repository interface {
	CreateSpot(spot *entities.CreateSpotInputSchema, ctx *fiber.Ctx) (*entities.Spot, error)
	ReadSpot() (*[]presenter.Spot, error)
	GetSpot(id int) (*entities.Spot, error)
	UpdateSpot(spot *entities.Spot, id int) (*entities.Spot, error)
	GetFindById(id int) (*entities.Spot, error)
	PartialUpdateSpot(spot *entities.Spot, id int) (*entities.Spot, error)
	DeleteSpot(id int) error

	// User
	GetUserById(id int) (*entities.User, error)
	validToken(t *jwt.Token, id string) bool
	getValueFromToken(key string, ctx *fiber.Ctx) int
}

type repository struct {
	DBConn *gorm.DB
}

func NewRepo(dbconn *gorm.DB) Repository {
	return &repository{
		DBConn: dbconn,
	}
}

// User Implement
func (r *repository) GetUserById(id int) (*entities.User, error) {
	var user entities.User
	if err := r.DBConn.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) getValueFromToken(key string, ctx *fiber.Ctx) int {
	token := ctx.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	value := int(claims[key].(float64))
	return value
}

func (r *repository) validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

// Spot
func (r *repository) GetSpot(id int) (*entities.Spot, error) {
	var spot entities.Spot
	result := r.DBConn.First(&spot, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &spot, nil
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
	userId := r.getValueFromToken("user_id", ctx)
	user, err := r.GetUserById(userId)
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
	//spot.User = *user

	result := r.DBConn.Create(&spot)
	if result.Error != nil {
		return nil, result.Error
	}
	return &spot, nil
}

func (r *repository) ReadSpot() (*[]presenter.Spot, error) {
	var spots []presenter.Spot
	result := r.DBConn.Find(&spots)
	if result.Error != nil {
		return nil, result.Error
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
