package entity

import "time"

type WishList struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(150)"`

	//Spots     interfaces.SpotCollection `gorm:"-"`                       // GORM에서 무시
	Spots []Spot `gorm:"many2many:wishlist_spot"` // 실제 DB 관계

	UserId int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
