package category

type Service interface {
	GetAllCategory() (*[]Category, error)
}

type Repository interface {
	GetAllCategory() (*[]Category, error)
	// FindCategoryById(id string) (Category, error)
	// InsertCaterory(category Category) error
	// UpdateCategory(category Category) error
}
