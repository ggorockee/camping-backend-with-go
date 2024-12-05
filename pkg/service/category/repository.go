package category

import "gorm.io/gorm"

type Repository interface {
	GetCategoryList()
	CreateCategory()
	GetCategory()
	UpdateCategory()
	DeleteCategory()
}

type repository struct {
	DBConn *gorm.DB
}

// CreateCategory implements Repository.
func (r *repository) CreateCategory() {
	panic("unimplemented")
}

// DeleteCategory implements Repository.
func (r *repository) DeleteCategory() {
	panic("unimplemented")
}

// GetCategory implements Repository.
func (r *repository) GetCategory() {
	panic("unimplemented")
}

// GetCategoryList implements Repository.
func (r *repository) GetCategoryList() {
	panic("unimplemented")
}

// UpdateCategory implements Repository.
func (r *repository) UpdateCategory() {
	panic("unimplemented")
}

func NewRepo(dbconn *gorm.DB) Repository {
	return &repository{
		DBConn: dbconn,
	}
}
