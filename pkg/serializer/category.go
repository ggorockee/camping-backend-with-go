package serializer

import (
	"camping-backend-with-go/internal/domain"
	"camping-backend-with-go/pkg/dto"
)

type CategorySerializer interface {
	ListSerialize() dto.CategoryListOut
	DetailSerialize() dto.CategoryDetailOut
}

type serializer struct {
	category *entities.Category
}

func (s *serializer) ListSerialize() dto.CategoryListOut {
	return dto.CategoryListOut{
		Id:   int(s.category.Id),
		Name: s.category.Name,
	}
}

func (s *serializer) DetailSerialize() dto.CategoryDetailOut {
	return dto.CategoryDetailOut{
		Id:        int(s.category.Id),
		Name:      s.category.Name,
		UpdatedAt: s.category.UpdatedAt,
		CreatedAt: s.category.CreatedAt,
	}
}

func NewCategorySerializer(category *entities.Category) CategorySerializer {
	return &serializer{category: category}
}
