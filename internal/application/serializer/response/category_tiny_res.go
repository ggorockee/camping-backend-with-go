package response

type CategoryTinyRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	// Time Logging
	//UpdatedAt time.Time `json:"updated_at"`
	//CreatedAt time.Time `json:"created_at"`
}
