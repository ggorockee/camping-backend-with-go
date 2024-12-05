package entities

import "time"

type Category struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex"`

	// Time Logging
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

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

func (c *Category) ListSerialize() CategoryListOut {
	return CategoryListOut{
		Id:   int(c.Id),
		Name: c.Name,
	}
}

func (c *Category) DetailSerialize() CategoryDetailOut {
	return CategoryDetailOut{
		Id:        int(c.Id),
		Name:      c.Name,
		UpdatedAt: c.UpdatedAt,
		CreatedAt: c.CreatedAt,
	}
}

type CreateCategoryInput struct {
	Name string `json:"name"`
}

type UpdateCategoryInput struct {
	Name string `json:"name"`
}
