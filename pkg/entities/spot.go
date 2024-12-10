package entities

import (
	"time"
)

type Spot struct {
	Id     uint `json:"id" gorm:"primaryKey"`
	UserId int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`

	Name        string `json:"name"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	PetFriendly bool   `json:"pet_friendly"`

	CategoryId *int     `gorm:"default:null" json:"category_id"` // CategoryId가 null일 수가 있음
	Category   Category `gorm:"foreignKey:CategoryId;constraint:OnDelete:SET NULL;"`

	Amenities []Amenity `gorm:"many2many:spot_amenities"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
