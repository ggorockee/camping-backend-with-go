package serializer

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"time"
)

type WishListSerializer interface {
	Serialize() []WishListRes
	WithSpotSerializer(s func([]entities.Spot) []dto.SpotListOut) WishListSerializer
	WithUserSerializer(s func(entities.User) dto.TinyUserOut) WishListSerializer
}

type WishListRes struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(150)"`

	Spots []dto.SpotListOut

	//User dto.TinyUserOut

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type wishList struct {
	wishList        *[]entities.WishList
	spotsSerializer func([]entities.Spot) []dto.SpotListOut
	userSerializer  func(user entities.User) dto.TinyUserOut
}

func (w *wishList) WithSpotSerializer(s func([]entities.Spot) []dto.SpotListOut) WishListSerializer {
	w.spotsSerializer = s
	return w
}

func (w *wishList) WithUserSerializer(s func(entities.User) dto.TinyUserOut) WishListSerializer {
	w.userSerializer = s
	return w
}

// Serialize implements WishListSerializer.
func (w *wishList) Serialize() []WishListRes {
	wishListsRes := make([]WishListRes, 0)

	// some logic
	for _, wl := range *w.wishList {
		wishItem := WishListRes{
			Id:        wl.Id,
			Name:      wl.Name,
			UpdatedAt: wl.UpdatedAt,
			CreatedAt: wl.CreatedAt,
		}

		if w.spotsSerializer != nil {
			wishItem.Spots = w.spotsSerializer(wl.Spots)
		}

		//if w.userSerializer != nil {
		//	wishItem.User = w.userSerializer(wl.User)
		//}

		wishListsRes = append(wishListsRes, wishItem)
	}

	return wishListsRes
}

func NewWishListSerializer(w *[]entities.WishList) WishListSerializer {
	return &wishList{wishList: w}
}
