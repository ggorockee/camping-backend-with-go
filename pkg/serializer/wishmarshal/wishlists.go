package wishmarshal

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
)

type WishListsSerializer interface {
	Serialize() []WishListRes
	WithSpotsSerializer(s func([]entities.Spot) []dto.SpotListOut) WishListsSerializer
	WithUserSerializer(s func(entities.User) dto.TinyUserOut) WishListsSerializer
}

type wishLists struct {
	wishLists       *[]entities.WishList
	spotsSerializer func(spots []entities.Spot) []dto.SpotListOut
	userSerializer  func(user entities.User) dto.TinyUserOut
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
			wishItem.Spots = w.spotsSerializer(wl.Spots.ToSlice())
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
