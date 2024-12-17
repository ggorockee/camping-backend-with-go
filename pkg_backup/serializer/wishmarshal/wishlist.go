package wishmarshal

import (
	entities2 "camping-backend-with-go/internal_backup/domain"
	"camping-backend-with-go/pkg/dto"
)

type WishListSerializer interface {
	Serialize() WishListRes
	WishSpotsSerializer(s func([]entities2.Spot) []dto.SpotListOut) WishListSerializer
}

type wishList struct {
	wishList        *entities2.WishList
	spotsSerializer func(spots []entities2.Spot) []dto.SpotListOut
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

func (w *wishList) WishSpotsSerializer(s func([]entities2.Spot) []dto.SpotListOut) WishListSerializer {
	w.spotsSerializer = s
	return w
}

func NewWishListSerializer(w *entities2.WishList) WishListSerializer {
	return &wishList{
		wishList: w,
	}
}
