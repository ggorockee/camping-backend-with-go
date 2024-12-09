package dto

import "time"

// ============= input schema =============

type CreateSpotIn struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	Review   string `json:"review"`
}

type UpdateSpotIn struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	Review   string `json:"review"`
}

// ============= output schema =============

type SpotListOut struct {
	Id        int         `json:"id"`
	User      TinyUserOut `json:"user"`
	Title     string      `json:"title"`
	Location  string      `json:"location"`
	Author    string      `json:"author"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Review    string      `json:"review"`
}

type SpotDetailOut struct {
	Id        int         `json:"id"`
	User      TinyUserOut `json:"user"`
	Title     string      `json:"title"`
	Location  string      `json:"location"`
	Author    string      `json:"author"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Review    string      `json:"review"`
}
