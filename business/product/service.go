package product

import (
	"AltaStore/business"
	"AltaStore/util/validator"
	"time"
)

type InsertProductSpec struct {
	UserId            string `validate:"required"`
	Code              string `validate:"required"`
	Name              string `validate:"required"`
	Price             int64  `validate:"required"`
	IsActive          bool
	ProductCategoryId string `validate:"required"`
	UnitName          string `validate:"required"`
	Description       string
}

type UpdateProductSpec struct {
	UserId            string `validate:"required"`
	Name              string `validate:"required"`
	Price             int64  `validate:"required"`
	IsActive          bool   `validate:"required"`
	ProductCategoryId string `validate:"required"`
	UnitName          string `validate:"required"`
	Description       string
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAllProduct() (*[]Product, error) {
	return s.repository.GetAllProduct()
}

func (s *service) GetAllProductByParameter(id, isActive, categoryName, code, name string) (*[]Product, error) {
	return s.repository.GetAllProductByParameter(id, isActive, categoryName, code, name)
}

func (s *service) FindProductById(id string) (*Product, error) {
	return s.repository.FindProductById(id)
}

func (s *service) InsertProduct(product *InsertProductSpec) error {
	err := validator.GetValidator().Struct(product)
	if err != nil {
		return business.ErrInvalidSpec
	}
	data := NewProduct(
		product.Code,
		product.Name,
		product.Price,
		product.IsActive,
		product.ProductCategoryId,
		product.UnitName,
		product.Description,
		product.UserId,
		time.Now(),
	)
	return s.repository.InsertProduct(data)
}

func (s *service) UpdateProduct(id string, updateProduct *UpdateProductSpec) error {
	product, err := s.repository.FindProductById(id)
	if err != nil {
		return err
	} else if product == nil {
		return business.ErrNotFound
	} else if product.DeletedBy != "" {
		return business.ErrDeleted
	}
	err = validator.GetValidator().Struct(product)
	if err != nil {
		return business.ErrInvalidSpec
	}
	modifiedproduct := product.ModifyProduct(
		updateProduct.Name,
		updateProduct.Price,
		updateProduct.IsActive,
		updateProduct.ProductCategoryId,
		updateProduct.UnitName,
		updateProduct.Description,
		updateProduct.UserId,
		time.Now())

	return s.repository.UpdateProduct(modifiedproduct)
}

func (s *service) DeleteProduct(id string, userid string) error {
	product, err := s.repository.FindProductById(id)
	if err != nil {
		return err
	} else if product == nil {
		return business.ErrNotFound
	} else if product.DeletedBy != "" {
		return business.ErrDeleted
	}

	deleteProduct := product.DeleteProduct(
		time.Now(),
		userid,
	)
	return s.repository.DeleteProduct(deleteProduct)
}
