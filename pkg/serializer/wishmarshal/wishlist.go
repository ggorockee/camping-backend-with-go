package wishmarshal

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
)

type WishListSerializer interface {
	Serialize() WishListRes
	WishSpotsSerializer(s func([]entities.Spot) []dto.SpotListOut) WishListSerializer
}

type wishList struct {
	wishList        *entities.WishList
	spotsSerializer func(spots []entities.Spot) []dto.SpotListOut
}

func (w *wishList) Serialize() WishListRes {
	wishListRes := WishListRes{
		Id:        w.wishList.Id,
		Name:      w.wishList.Name,
		UpdatedAt: w.wishList.UpdatedAt,
		CreatedAt: w.wishList.CreatedAt,
	}

	if w.spotsSerializer != nil {
		wishListRes.Spots = w.spotsSerializer(w.wishList.Spots)
	}

	return wishListRes
}

func (w *wishList) WishSpotsSerializer(s func([]entities.Spot) []dto.SpotListOut) WishListSerializer {
	w.spotsSerializer = s
	return w
}

func NewWishListSerializer(w *entities.WishList) WishListSerializer {
	return &wishList{
		wishList: w,
	}
}
