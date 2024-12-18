package response

import (
	"time"
)

type SpotDetailRes struct {
	Id uint `json:"id"`

	User UserTinyRes `json:"user"`

	Name        string `json:"name"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	PetFriendly bool   `json:"pet_friendly"`

	Category CategoryTinyRes `json:"category"`

	Amenities []AmenityTinyRes `gorm:"many2many:spot_amenities"`

	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Reviews   []ReviewTinyRes `gorm:"foreignKey:SpotId" json:"reviews"`
}
