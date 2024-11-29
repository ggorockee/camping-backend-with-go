package entities

import (
	"time"
)

type Spot struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteRequest struct {
	Id string `json:"id"`
}

type CreateSpotInputSchema struct {
	Title    string `json:"title"`
	Location string `json:"location"`
}

type UpdateSpotSchema struct {
	Title    string `json:"title"`
	Location string `json:"location"`
}
