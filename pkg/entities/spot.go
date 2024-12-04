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
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Spot) ListSerialize() SpotListOutputSchema {
	return SpotListOutputSchema{
		Id:        s.Id,
		User:      s.User.TinyUserSerialize(),
		Title:     s.Title,
		Location:  s.Location,
		Author:    s.Author,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		Review:    s.Review,
	}
}

func (s *Spot) DetailSerialize() SpotDetailOutputSchema {
	return SpotDetailOutputSchema{
		Id:        s.Id,
		User:      s.User.TinyUserSerialize(),
		Title:     s.Title,
		Location:  s.Location,
		Author:    s.Author,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
		Review:    s.Review,
	}
}

type DeleteRequest struct {
	Id string `json:"id"`
}

type CreateSpotInputSchema struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	Review   string `json:"review"`
}

type UpdateSpotSchema struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	Review   string `json:"review"`
}

type SpotListOutputSchema struct {
	Id        uint                 `json:"id"`
	User      TinyUserOutputSchema `json:"user"`
	Title     string               `json:"title"`
	Location  string               `json:"location"`
	Author    string               `json:"author"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	Review    string               `json:"review"`
}

type SpotDetailOutputSchema struct {
	Id        uint                 `json:"id"`
	User      TinyUserOutputSchema `json:"user"`
	Title     string               `json:"title"`
	Location  string               `json:"location"`
	Author    string               `json:"author"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	Review    string               `json:"review"`
}
