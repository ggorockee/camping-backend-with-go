package entities

import "time"

type WishList struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(150)"`

	Spots []Spot `gorm:"many2many:wishlist_spot"`

	UserId int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
