package wishlist

import (
	wishlistdto "camping-backend-with-go/internal/application/dto/wishlist"
	"camping-backend-with-go/internal/domain/entity"
	spotrepository "camping-backend-with-go/internal/domain/repository/spot"
	"camping-backend-with-go/pkg/util"

	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetWishLists(context ...*fiber.Ctx) (*[]entity.WishList, error)
	CreateWishList(input *wishlistdto.CreateWishListReq, context ...*fiber.Ctx) (*entity.WishList, error)
	GetWishList(id int, context ...*fiber.Ctx) (*entity.WishList, error)
	UpdateWishList(input *wishlistdto.UpdateWishListReq, id int, context ...*fiber.Ctx) (*entity.WishList, error)
	DeleteWishList(id int, context ...*fiber.Ctx) error
	WishListToggle(wishListId int, spotId int, context ...*fiber.Ctx) error
}

type repository struct {
	dbConn   *gorm.DB
	spotRepo spotrepository.SpotRepository
}

func (r *repository) GetWishLists(context ...*fiber.Ctx) (*[]entity.WishList, error) {
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	requestUser, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return nil, errors.New("user is not authenticated")
	}

	var wishLists []entity.WishList
	if err := r.dbConn.
		Preload("User").
		Preload("Spots").
		Where("user_id = ?", requestUser.Id).
		Find(&wishLists).Error; err != nil {
		return nil, err
	}

	return &wishLists, err
}

func (r *repository) CreateWishList(input *wishlistdto.CreateWishListReq, context ...*fiber.Ctx) (*entity.WishList, error) {
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	requestUser, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return nil, errors.New("user is not authenticated")
	}

	wishlist := entity.WishList{
		Name: input.Name,
		//Spots:     nil,
		User:      requestUser,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	if err := r.dbConn.Create(&wishlist).Error; err != nil {
		return nil, err
	}

	return &wishlist, nil
}

func (r *repository) GetWishList(id int, context ...*fiber.Ctx) (*entity.WishList, error) {
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	requestUser, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return nil, errors.New("user is not authenticated")
	}

	var wishList entity.WishList

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

func (r *repository) UpdateWishList(input *wishlistdto.UpdateWishListReq, id int, context ...*fiber.Ctx) (*entity.WishList, error) {
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	requestUser, ok := c.Locals("request_user").(entity.User)
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

func (r *repository) DeleteWishList(id int, context ...*fiber.Ctx) error {
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	wishList, err := r.GetWishList(id, c)
	if err != nil {
		return err
	}

	requestUser, ok := c.Locals("request_user").(entity.User)
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

func (r *repository) WishListToggle(wishListId int, spotId int, context ...*fiber.Ctx) error {
	//c, err := util.ContextParser(context...)
	//util.HandleFunc(err)

	//wishList, err := r.GetWishList(wishListId, c)
	//if err != nil {
	//	return err
	//}
	//
	//spot, err := r.spotRepo.GetSpotById(spotId)
	//if err != nil {
	//	return err
	//}
	//
	////if wishList.Spots.Filter(int(spot.Id)).Exists() {
	////	wishList.Spots.Remove(spot)
	////} else {
	////	wishList.Spots.Add(spot)
	////}

	return nil
}

func NewRepo(dbConn *gorm.DB, s spotrepository.SpotRepository) Repository {
	return &repository{
		dbConn:   dbConn,
		spotRepo: s,
	}
}
