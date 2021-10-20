package category

import (
	"AltaStore/business/category"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type ProductCategory struct {
	ID        string    `gorm:"id;type:uuid;primaryKey"`
	Code      string    `gorm:"code;type:varchar(50);unique"`
	Name      string    `gorm:"name;type:varchar(100)"`
	CreatedAt time.Time `gorm:"created_at"`
	CreatedBy string    `gorm:"created_by;type:uuid"`
	UpdatedAt time.Time `gorm:"updated_at"`
	UpdatedBy string    `gorm:"updated_by;type:uuid"`
	DeletedAt time.Time `gorm:"deleted_at"`
	DeletedBy string    `gorm:"deleted_by;type:uuid"`
}

func (p *ProductCategory) toCategory() category.Category {
	return category.Category{
		ID:        p.ID,
		Code:      p.Code,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		CreatedBy: p.CreatedBy,
		UpdatedAt: p.UpdatedAt,
		UpdatedBy: p.UpdatedBy,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllCategory() (*[]category.Category, error) {
	var categories []ProductCategory

	err := r.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	var result []category.Category
	for _, value := range categories {
		result = append(result, value.toCategory())
	}

	return &result, nil
}

// func (r *Repository) FindCategoryById(id string) (*category.Category, error) {

// }

// func (r *Repository) InsertCategory(category category.Category) error {

// }

// func (r *Repository) UpdateCategory(category category.Category) error {

// }
