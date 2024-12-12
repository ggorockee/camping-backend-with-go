package serializer

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"time"
)

type WishListsSerializer interface {
	Serialize() []WishListRes
	WithSpotsSerializer(s func([]entities.Spot) []dto.SpotListOut) WishListsSerializer
	WithUserSerializer(s func(entities.User) dto.TinyUserOut) WishListsSerializer
}

type WishListSerializer interface {
	Serialize() WishListRes
	WishSpotsSerializer(s func([]entities.Spot) []dto.SpotListOut) WishListSerializer
}

type WishListRes struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(150)"`

	Spots []dto.SpotListOut

	//User dto.TinyUserOut

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type wishLists struct {
	wishLists       *[]entities.WishList
	spotsSerializer func(spots []entities.Spot) []dto.SpotListOut
	userSerializer  func(user entities.User) dto.TinyUserOut
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

func (w *wishLists) WithSpotsSerializer(s func([]entities.Spot) []dto.SpotListOut) WishListsSerializer {
	w.spotsSerializer = s
	return w
}

func (w *wishLists) WithUserSerializer(s func(entities.User) dto.TinyUserOut) WishListsSerializer {
	w.userSerializer = s
	return w
}

// Serialize implements WishListSerializer.
func (w *wishLists) Serialize() []WishListRes {
	wishListsRes := make([]WishListRes, 0)

	// some logic
	for _, wl := range *w.wishLists {
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

func NewWishListsSerializer(w *[]entities.WishList) WishListsSerializer {
	return &wishLists{wishLists: w}
}

func NewWishListSerializer(w *entities.WishList) WishListSerializer {
	return &wishList{
		wishList: w,
	}
}
