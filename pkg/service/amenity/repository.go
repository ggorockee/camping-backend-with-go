package amenity

import (
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	Create(input *entities.CreateAmenityInput, ctx *fiber.Ctx) (*entities.Amenity, error)
}

type repository struct {
	DBConn   *gorm.DB
	UserRepo user.Repository
}

func (r *repository) Create(input *entities.CreateAmenityInput, ctx *fiber.Ctx) (*entities.Amenity, error) {
	_ = r.UserRepo.GetValueFromToken("user_id", ctx)

	var amenity entities.Amenity

	amenity.Name = input.Name
	amenity.Description = input.Description
	amenity.CreatedAt = time.Now()
	amenity.UpdatedAt = time.Now()

	if err := r.DBConn.Create(&amenity).Error; err != nil {
		return nil, err
	}

	return &amenity, nil
}

func NewRepo(dbConn *gorm.DB, userRepo user.Repository) Repository {
	return &repository{DBConn: dbConn, UserRepo: userRepo}
}