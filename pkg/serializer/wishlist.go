package serializer

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"time"
)

type WishListSerializer interface {
	Serialize() []WishListRes
}

type WishListRes struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(150)"`

	Spots []dto.SpotListOut

	User dto.TinyUserOut

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type wishList struct {
	wishList    *[]entities.WishList
	serializers any
}

// Serialize implements WishListSerializer.
func (w *wishList) Serialize() []WishListRes {
	var wishListsRes []WishListRes

	// some logic

	return wishListsRes
}

func NewWishListSerializer(w *[]entities.WishList, serializers ...any) WishListSerializer {
	return &wishList{wishList: w, serializers: serializers}
}
