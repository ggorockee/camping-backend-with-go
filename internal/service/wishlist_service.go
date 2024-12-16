package service

import (
	"camping-backend-with-go/internal/domain"
	repository2 "camping-backend-with-go/internal/repository"
	"camping-backend-with-go/pkg/dto/wishdto"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetWishLists(contexts ...*fiber.Ctx) (*[]domain.WishList, error)
	WishListToggle(wishListId int, spotId int, contexts ...*fiber.Ctx) error
	CreateWishList(input *wishdto.CreateWishListReq, contexts ...*fiber.Ctx) (*domain.WishList, error)
	GetWishList(id int, contexts ...*fiber.Ctx) (*domain.WishList, error)
	UpdateWishList(input *wishdto.UpdateWishListReq, id int, contexts ...*fiber.Ctx) (*domain.WishList, error)
	DeleteWishList(id int, contexts ...*fiber.Ctx) error
}

type service struct {
	repository repository2.Repository
}

func (s *service) CreateWishList(input *wishdto.CreateWishListReq, contexts ...*fiber.Ctx) (*domain.WishList, error) {
	return s.repository.CreateWishList(input, contexts...)
}

func (s *service) GetWishList(id int, contexts ...*fiber.Ctx) (*domain.WishList, error) {
	return s.repository.GetWishList(id, contexts...)
}

func (s *service) UpdateWishList(input *wishdto.UpdateWishListReq, id int, contexts ...*fiber.Ctx) (*domain.WishList, error) {
	return s.repository.UpdateWishList(input, id, contexts...)
}

func (s *service) DeleteWishList(id int, contexts ...*fiber.Ctx) error {
	return s.repository.DeleteWishList(id, contexts...)
}

func (s *service) WishListToggle(wishListId int, spotId int, contexts ...*fiber.Ctx) error {
	return s.repository.WishListToggle(wishListId, spotId, contexts...)
}

func (s *service) GetWishLists(contexts ...*fiber.Ctx) (*[]domain.WishList, error) {
	return s.repository.GetWishLists(contexts...)
}

func NewService(r repository2.Repository) Service {
	return &service{repository: r}
}
