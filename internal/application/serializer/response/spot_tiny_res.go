package response

import "time"

type SpotTinyRes struct {
	Id   uint        `json:"id" gorm:"primaryKey"`
	User UserTinyRes `json:"user"`

	Name    string `json:"name"`
	Country string `json:"country"`
	City    string `json:"city"`
	Price   int    `json:"price"`
	Address string `json:"address"`

	Category CategoryTinyRes `json:"category"`

	Amenities []AmenityTinyRes `json:"amenities"`
	Reviews   []ReviewTinyRes  `json:"reviews"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
