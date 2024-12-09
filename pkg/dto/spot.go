package dto

import "time"

// ============= input schema =============

type CreateSpotIn struct {
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	Price       int     `json:"price"`
	Description *string `json:"description"`
	Address     string  `json:"address"`
	PetFriendly bool    `json:"pet_friendly"`
	Category    int     `json:"category"`
	Amenities   *[]int  `json:"amenities"`
}

type UpdateSpotIn struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	Review   string `json:"review"`
}

// ============= output schema =============

type SpotListOut struct {
	Id          int               `json:"id"`
	User        TinyUserOut       `json:"user"`
	Name        string            `json:"name"`
	Country     string            `json:"country"`
	City        string            `json:"city"`
	Price       int               `json:"price"`
	Description *string           `json:"description"`
	Address     string            `json:"address"`
	PetFriendly bool              `json:"pet_friendly"`
	Category    int               `json:"category"`
	Amenities   *[]AmenityListOut `json:"amenities"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type SpotDetailOut struct {
	Id          int              `json:"id"`
	User        TinyUserOut      `json:"user"`
	Name        string           `json:"name"`
	Country     string           `json:"country"`
	City        string           `json:"city"`
	Price       int              `json:"price"`
	Description string           `json:"description"`
	Address     string           `json:"address"`
	PetFriendly bool             `json:"pet_friendly"`
	Category    int              `json:"category"`
	Amenities   []AmenityListOut `json:"amenities"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}
