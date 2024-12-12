package serializer

import (
	"camping-backend-with-go/pkg/entities"
	"time"
)

type WishListSerializer interface {
	Serialize() []WishListRes
}

type WishListRes struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(150)"`

	Spots []entities.Spot `gorm:"many2many:wishlist_spot"`

	UserId int           `json:"user_id"`
	User   entities.User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type wishlist struct {
	wishlist *[]entities.WishList
}

// Serialize implements WishListSerializer.
func (w *wishlist) Serialize() []WishListRes {
	var wishListRes []WishListRes
	for _, wishItem := range *w.wishlist {

	}
	return WishListRes{
		Id:        w.wishlist.Id,
		Name:      w.wishlist.Name,
		Spots:     w.wishlist.Spots,
		UserId:    w.wishlist.UserId,
		User:      w.wishlist.User,
		UpdatedAt: w.wishlist.UpdatedAt,
		CreatedAt: w.wishlist.CreatedAt,
	}
}

func NewWishListSerializer(w *[]entities.WishList) WishListSerializer {
	return &wishlist{wishlist: w}
}
