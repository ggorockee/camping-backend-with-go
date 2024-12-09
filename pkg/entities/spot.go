package entities

import (
	"time"
)

type Spot struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	UserId     uint      `json:"user_id"`
	User       User      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	CategoryId *int      `gorm:"default:null" json:"category_id"` // CategoryId가 null일 수가 있음
	Category   Category  `gorm:"foreignKey:CategoryId;constraint:OnDelete:SET NULL;"`
	Title      string    `json:"title"`
	Location   string    `json:"location"`
	Author     string    `json:"author"`
	Review     string    `json:"review"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	CoverImg string `json:"cover_img"`
}
