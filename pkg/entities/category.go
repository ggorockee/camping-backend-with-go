package entities

import "time"

type Category struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"uniqueIndex"`
	CoverImg string `json:"cover_img"`

	// Time Logging
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
