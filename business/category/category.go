package category

import "time"

// Membuat core data pada domain bisnis
type Category struct {
	ID        string
	Code      string
	Name      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

// fungsi menambahkan data kategory
func NewProductCategory(
	id string, code string, name string,
	creator string, createdAt time.Time,
) Category {
	return Category{
		ID:        id,
		Code:      code,
		Name:      name,
		CreatedAt: createdAt,
		CreatedBy: creator,
		UpdatedAt: createdAt,
		UpdatedBy: creator,
	}
}

// fungsi memperbarui data
func (p *Category) UpdateProductCategory(code string,
	name string, updater string, updatedAt time.Time,
) Category {
	return Category{
		ID:        p.ID,
		Code:      code,
		Name:      name,
		CreatedAt: p.CreatedAt,
		CreatedBy: p.CreatedBy,
		UpdatedAt: updatedAt,
		UpdatedBy: updater,
	}
}
