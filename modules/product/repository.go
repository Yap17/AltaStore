package product

import (
	"AltaStore/business/product"
	category "AltaStore/modules/category"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Product struct {
	ID                string `gorm:"id;type:uuid;primaryKey"`
	Code              string `gorm:"code;type:varchar(50);unique"`
	Name              string `gorm:"name;type:varchar(100)"`
	Price             int64  `gorm:"price"`
	ProductCategoryId string `gorm:"productcategoryid;type:uuid"`
	ProductCategory   category.ProductCategory
	IsActive          bool      `gorm:"isactive;type:boolean"`
	UnitName          string    `gorm:"unitname;type:varchar(50)"`
	Description       string    `gorm:"description;type:varchar(255)"`
	CreatedAt         time.Time `gorm:"created_at"`
	CreatedBy         string    `gorm:"created_by;type:varchar(50)"`
	UpdatedAt         time.Time `gorm:"updated_at"`
	UpdatedBy         string    `gorm:"updated_by;type:varchar(50)"`
	DeletedAt         time.Time `gorm:"deleted_at"`
	DeletedBy         string    `gorm:"deleted_by;type:varchar(50)"`
}

func (p *Product) ToProduct() product.Product {
	return product.Product{
		ID:                p.ID,
		Code:              p.Code,
		Name:              p.Name,
		Price:             p.Price,
		IsActive:          p.IsActive,
		ProductCategoryId: p.ProductCategoryId,
		UnitName:          p.UnitName,
		Description:       p.Description,
		CreatedAt:         p.CreatedAt,
		CreatedBy:         p.CreatedBy,
		UpdatedAt:         p.UpdatedAt,
		UpdatedBy:         p.UpdatedBy,
		DeletedAt:         p.DeletedAt,
		DeletedBy:         p.DeletedBy,
	}
}

func newDataProduct(p product.Product) *Product {
	return &Product{
		ID:                p.ID,
		Code:              p.Code,
		Name:              p.Name,
		Price:             p.Price,
		IsActive:          p.IsActive,
		ProductCategoryId: p.ProductCategoryId,
		UnitName:          p.UnitName,
		Description:       p.Description,
		CreatedAt:         p.CreatedAt,
		CreatedBy:         p.CreatedBy,
		UpdatedAt:         p.CreatedAt,
		UpdatedBy:         p.CreatedBy,
		DeletedAt:         p.DeletedAt,
		DeletedBy:         p.DeletedBy,
	}
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllProduct() (*[]product.Product, error) {
	var products []Product
	var tempProduct product.Product

	err := r.DB.Preload("ProductCategory").Where("deleted_by = ''").Find(&products).Error
	if err != nil {
		return nil, err
	}

	var productOuts []product.Product
	for _, value := range products {
		tempProduct = value.ToProduct()
		tempProduct.ProductCategoryName = value.ProductCategory.Name
		tempProduct.Qty = 10
		tempProduct.QtyCart = 0
		productOuts = append(productOuts, tempProduct)
	}

	return &productOuts, nil
}

func (r *Repository) GetAllProductByParameter(id, isActive, categoryName,
	code, name string) (*[]product.Product, error) {
	var products []Product
	var tempProduct product.Product
	var productOuts []product.Product

	temp := r.DB
	temp = temp.Preload("ProductCategory")
	if categoryName != "" {
		temp = temp.Joins("inner join product_categories on product_categories.id = products.product_category_id")
		temp = temp.Where("product_categories.name = ?", categoryName)
	}
	if id != "" {
		temp = temp.Where("id = ?", id)
	}
	if isActive != "" {
		res, err := strconv.ParseBool(isActive)
		if err != nil {
			return &productOuts, nil
		}
		temp = temp.Where("is_active = ?", res)
	}
	if code != "" {
		temp = temp.Where("code = ?", code)
	}
	if name != "" {
		temp = temp.Where("name = ?", name)
	}

	err := temp.Where("deleted_by = ''").Find(&products).Error
	if err != nil {
		return nil, err
	}

	for _, value := range products {
		tempProduct = value.ToProduct()

		tempProduct.ProductCategoryName = value.ProductCategory.Name
		tempProduct.Qty = 10
		tempProduct.QtyCart = 0
		productOuts = append(productOuts, tempProduct)
	}

	return &productOuts, nil
}

func (r *Repository) FindProductById(id string) (*product.Product, error) {
	var product *Product

	err := r.DB.Where("id = ?", id).Where("deleted_by = ''").First(&product).Error
	if err != nil {
		return nil, err
	}

	response := product.ToProduct()
	return &response, nil
}

func (r *Repository) InsertProduct(p product.Product) error {
	product := newDataProduct(p)
	if err := r.DB.Create(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateProduct(p product.Product) error {
	product := newDataProduct(p)
	err := r.DB.Model(&product).Updates(Product{
		Name:              product.Name,
		Price:             product.Price,
		IsActive:          product.IsActive,
		ProductCategoryId: product.ProductCategoryId,
		UnitName:          product.UnitName,
		Description:       product.Description,
		UpdatedAt:         product.UpdatedAt,
		UpdatedBy:         product.UpdatedBy,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteProduct(p product.Product) error {
	product := newDataProduct(p)

	err := r.DB.Model(&product).Updates(Product{
		DeletedBy: product.DeletedBy,
		DeletedAt: product.DeletedAt,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
