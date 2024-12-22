package spotdto

type CreateSpotReq struct {
	Name        *string   `json:"name"`
	Country     *string   `json:"country"`
	City        *string   `json:"city"`
	Price       *int      `json:"price"`
	Description *string   `json:"description"`
	Address     *string   `json:"address"`
	PetFriendly *bool     `json:"pet_friendly"`
	Category    *string   `json:"category"`
	Amenities   *[]string `json:"amenities"`
}
