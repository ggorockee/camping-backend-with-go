package repository

import (
	entities2 "camping-backend-with-go/internal_backup/domain"
	"camping-backend-with-go/pkg/dto/wishdto"
	"camping-backend-with-go/pkg/service/spot"
	"camping-backend-with-go/pkg/service/util"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository interface {
	GetWishLists(contexts ...*fiber.Ctx) (*[]entities2.WishList, error)
	CreateWishList(input *wishdto.CreateWishListReq, contexts ...*fiber.Ctx) (*entities2.WishList, error)
	GetWishList(id int, contexts ...*fiber.Ctx) (*entities2.WishList, error)
	UpdateWishList(input *wishdto.UpdateWishListReq, id int, contexts ...*fiber.Ctx) (*entities2.WishList, error)
	DeleteWishList(id int, contexts ...*fiber.Ctx) error
	WishListToggle(wishListId int, spotId int, contexts ...*fiber.Ctx) error
}

type repository struct {
	dbConn   *gorm.DB
	spotRepo spot.Repository
}

func (r *repository) WishListToggle(wishListId int, spotId int, contexts ...*fiber.Ctx) error {
	c, err := util.ContextParser(contexts...)
	if err != nil {
		return err
	}

	fetchedWishList, err := r.GetWishList(wishListId, c)
	if err != nil {
		return err
	}

	fetchedSpot, err := r.spotRepo.GetFindById(spotId)
	if err != nil {
		return err
	}

	if fetchedWishList.Spots.Filter(int(fetchedSpot.Id)).Exists() {
		fetchedWishList.Spots.Remove(fetchedSpot)
	} else {
		fetchedWishList.Spots.Add(fetchedSpot)
	}

	return nil
}

func (r *repository) CreateWishList(input *wishdto.CreateWishListReq, contexts ...*fiber.Ctx) (*entities2.WishList, error) {
	c, err := util.ContextParser(contexts...)
	if err != nil {
		return nil, err
	}

	requestUser, ok := c.Locals("request_user").(entities2.User)
	if !ok {
		return nil, errors.New("user is not authenticated")
	}

	wishlist := entities2.WishList{
		Name:      input.Name,
		Spots:     nil,
		User:      requestUser,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	if err := r.dbConn.Create(&wishlist).Error; err != nil {
		return nil, err
	}

	return &wishlist, nil
}

func (r *repository) GetWishList(id int, contexts ...*fiber.Ctx) (*entities2.WishList, error) {
	c, err := util.ContextParser(contexts...)
	if err != nil {
		return nil, err
	}

	requestUser, ok := c.Locals("request_user").(entities2.User)
	if !ok {
		return nil, errors.New("user is not authenticated")
	}

	var wishList entities2.WishList

	if err := r.dbConn.
		Preload("User").
		Preload("Spots").
		Where("user_id = ?", id).
		Where("id = ?", requestUser.Id).
		Find(&wishList).Error; err != nil {
		return nil, err
	}

	return &wishList, nil
}

func (r *repository) UpdateWishList(input *wishdto.UpdateWishListReq, id int, contexts ...*fiber.Ctx) (*entities2.WishList, error) {
	c, err := util.ContextParser(contexts...)
	if err != nil {
		return nil, err
	}

	requestUser, ok := c.Locals("request_user").(entities2.User)
	if !ok {
		return nil, errors.New("user is not authenticated")
	}

	wishList, err := r.GetWishList(id, c)
	if err != nil {
		return nil, err
	}

	if wishList.UserId != int(requestUser.Id) {
		return nil, errors.New("permission denied")
	}

	if input.Name != "" {
		wishList.Name = input.Name
	}

	if err := r.dbConn.Model(wishList).Updates(wishList).Error; err != nil {
		return nil, err
	}

	return wishList, nil

}

func (r *repository) DeleteWishList(id int, contexts ...*fiber.Ctx) error {
	c, err := util.ContextParser(contexts...)
	if err != nil {
		return err
	}

	wishList, err := r.GetWishList(id, c)
	if err != nil {
		return err
	}

	requestUser, ok := c.Locals("request_user").(entities2.User)
	if !ok {
		return errors.New("user is not authenticated")
	}

	if wishList.UserId != int(requestUser.Id) {
		return errors.New("permission denied")
	}

	if err := r.dbConn.Delete(wishList).Error; err != nil {
		return err
	}

	return nil
}

// GetWishList implements Repository.
func (r *repository) GetWishLists(contexts ...*fiber.Ctx) (*[]entities2.WishList, error) {
	c, err := util.ContextParser(contexts...)
	if err != nil {
		return nil, err
	}

	requestUser, ok := c.Locals("request_user").(entities2.User)
	if !ok {
		return nil, errors.New("user is not authenticated")
	}

	var wishList []entities2.WishList
	if err := r.dbConn.
		Preload("User").
		Preload("Spots").
		Where("user_id = ?", requestUser.Id).
		Find(&wishList).Error; err != nil {
		return nil, err
	}

	return &wishList, err

}

func NewRepo(dbConn *gorm.DB, spotRepo spot.Repository) Repository {
	return &repository{
		dbConn:   dbConn,
		spotRepo: spotRepo,
	}
}
