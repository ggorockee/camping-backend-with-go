package handlers

import (
	"camping-backend-with-go/api/presenter"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/spot"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateSpot(spot *entities.Spot) (*entities.Spot, error)
	ReadSpot() (*[]presenter.Spot, error)
	UpdateSpot(spot *entities.Spot) (*entities.Spot, error)
	DeleteSpot(ID string) error
}

type repository struct {
	DBConn *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewRepo(dbConn *gorm.DB) Repository {
	return &repository{
		DBConn: dbConn,
	}
}

func AddSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(presenter.SpotSuccessResponse())
	}
}

func UpdateSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(presenter.SpotSuccessResponse())
	}
}

func RemoveSpot(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(presenter.SpotSuccessResponse())
	}
}

func GetSpots(service spot.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(presenter.SpotSuccessResponse())
	}
}
