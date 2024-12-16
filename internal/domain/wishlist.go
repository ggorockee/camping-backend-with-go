package domain

import (
	"camping-backend-with-go/pkg/interfaces"
	"fmt"
	"time"
)

type WishList struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(150)"`

	Spots     interfaces.SpotCollection `gorm:"-"`                       // GORM에서 무시
	spotsData []Spot                    `gorm:"many2many:wishlist_spot"` // 실제 DB 관계

	UserId int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

var NewSpotCollection interfaces.SpotCollectionFactory

// 도메인 Spot을 인터페이스 Spot으로 변환하는 함수 정의
func convertToInterfaceSpots(domainSpots []Spot) []interfaces.Spot {
	interfaceSpots := make([]interfaces.Spot, len(domainSpots))
	for i, spot := range domainSpots {
		interfaceSpots[i] = &spot
	}
	return interfaceSpots
}

func (w *WishList) AfterFind() error {
	// NewSpotCollection
	// []interfaces.Spot -> interfaces.SpotCollection

	// domainToInterface
	// []domain.spot ->

	// w.Spot 는 []domain.spot 이므로
	// []domain.spot 을 -> []interfaces.Spot 으로 만들어줄 장치가 필요함
	w.Spots = NewSpotCollection(convertToInterfaceSpots(w.spotsData))
	return nil
}

func (w *WishList) BeforeSave() error {
	if w.Spots != nil {
		interfacesSpots := w.Spots.ToSlice()
		w.spotsData = make([]Spot, len(interfacesSpots))
		for i, spot := range interfacesSpots {
			if concreteSpot, ok := spot.(*Spot); ok {
				w.spotsData[i] = *concreteSpot
			} else {
				return fmt.Errorf("invalid spot type at index %d", i)
			}
		}
	}

	return nil
}
