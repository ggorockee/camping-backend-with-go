package wishmarshal

import (
	"camping-backend-with-go/pkg/dto"
	"time"
)

type WishListRes struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(150)"`

	Spots []dto.SpotListOut

	//User dto.TinyUserOut

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
