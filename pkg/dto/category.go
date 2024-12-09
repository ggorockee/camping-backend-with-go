package dto

import "time"

//
// ========= Input Schema ============
//

type CreateCategoryIn struct {
	Name string `json:"name"`
}

type UpdateCategoryIn struct {
	Name string `json:"name"`
}

//
// ========== OutPut Schema ============
//

type CategoryListOut struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryDetailOut struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
