package wishlistsvc

import (
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/util"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository interface {
	GetWishList(contexts ...*fiber.Ctx) (*[]entities.WishList, error)
}

type repository struct {
	dbConn *gorm.DB
}

// GetWishList implements Repository.
func (r *repository) GetWishList(contexts ...*fiber.Ctx) (*[]entities.WishList, error) {
	c, err := util.ContextParser(contexts...)
	if err != nil {
		return nil, err
	}

	requestUser, ok := c.Locals("request_user").(entities.User)
	if !ok {
		return nil, errors.New("user is not authenticated")
	}

	var wishList []entities.WishList
	if err := r.dbConn.
		Preload("User").
		Preload("Spots").
		Where("user_id = ?", requestUser.Id).
		Find(&wishList).Error; err != nil {
		return nil, err
	}

	return &wishList, err

}

func NewRepo(dbConn *gorm.DB) Repository {
	return &repository{
		dbConn: dbConn,
	}
}
