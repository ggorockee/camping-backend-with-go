package entities

import "time"

type Category struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex"`

	// Time Logging
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type CategorySerializer interface {
	ListSerialize() CategoryListOut
	DetailSerialize() CategoryDetailOut
}

type serializer struct {
	category *Category
}

func (s *serializer) ListSerialize() CategoryListOut {
	return CategoryListOut{
		Id:   int(s.category.Id),
		Name: s.category.Name,
	}
}

func (s *serializer) DetailSerialize() CategoryDetailOut {
	return CategoryDetailOut{
		Id:        int(s.category.Id),
		Name:      s.category.Name,
		UpdatedAt: s.category.UpdatedAt,
		CreatedAt: s.category.CreatedAt,
	}
}

func NewCategorySerializer(category *Category) CategorySerializer {
	return &serializer{category: category}
}

//
// ========= Input Schema ============
//

type CreateCategoryInput struct {
	Name string `json:"name"`
}

type UpdateCategoryInput struct {
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
