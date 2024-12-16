package domain

import "time"

type Review struct {
	Id     uint `json:"id" gorm:"primaryKey"`
	UserId int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`

	SpotId *int `json:"spot_id"`
	Spot   Spot `gorm:"foreignKey:SpotId;constraint:OnDelete:SET NULL"`

	Payload string `json:"payload"`
	Rating  int    `json:"rating"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
