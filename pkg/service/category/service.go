package category

type Service interface {
	GetCategoryList()
	CreateCategory()
	GetCategory()
	UpdateCategory()
	DeleteCategory()
}

type service struct {
	repo Repository
}

// CreateCategory implements Service.
func (s *service) CreateCategory() {
	panic("unimplemented")
}

// DeleteCategory implements Service.
func (s *service) DeleteCategory() {
	panic("unimplemented")
}

// GetCategory implements Service.
func (s *service) GetCategory() {
	panic("unimplemented")
}

// GetCategoryList implements Service.
func (s *service) GetCategoryList() {
	panic("unimplemented")
}

// UpdateCategory implements Service.
func (s *service) UpdateCategory() {
	panic("unimplemented")
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
