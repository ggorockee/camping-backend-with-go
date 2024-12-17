package wishmarshal

import (
	entities2 "camping-backend-with-go/internal_backup/domain"
	"camping-backend-with-go/pkg/dto"
)

type WishListsSerializer interface {
	Serialize() []WishListRes
	WithSpotsSerializer(s func([]entities2.Spot) []dto.SpotListOut) WishListsSerializer
	WithUserSerializer(s func(entities2.User) dto.TinyUserOut) WishListsSerializer
}

type wishLists struct {
	wishLists       *[]entities2.WishList
	spotsSerializer func(spots []entities2.Spot) []dto.SpotListOut
	userSerializer  func(user entities2.User) dto.TinyUserOut
}

func (w *wishLists) WithSpotsSerializer(s func([]entities2.Spot) []dto.SpotListOut) WishListsSerializer {
	w.spotsSerializer = s
	return w
}

func (w *wishLists) WithUserSerializer(s func(entities2.User) dto.TinyUserOut) WishListsSerializer {
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
			wishItem.Spots = w.spotsSerializer(wl.Spots.ToSlice())
		}

		//if w.userSerializer != nil {
		//	wishItem.User = w.userSerializer(wl.User)
		//}

		wishListsRes = append(wishListsRes, wishItem)
	}

	return wishListsRes
}

func NewWishListsSerializer(w *[]entities2.WishList) WishListsSerializer {
	return &wishLists{wishLists: w}
}
