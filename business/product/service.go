package product

import (
	"AltaStore/business"
	"AltaStore/business/admin"
	"AltaStore/util/validator"
	"time"
)

type InsertProductSpec struct {
	AdminId           string `validate:"required"`
	Code              string `validate:"required"`
	Name              string `validate:"required"`
	Price             int64  `validate:"required"`
	IsActive          bool
	ProductCategoryId string `validate:"required"`
	UnitName          string `validate:"required"`
	Description       string
}

type UpdateProductSpec struct {
	AdminId           string `validate:"required"`
	Name              string `validate:"required"`
	Price             int64  `validate:"required"`
	IsActive          bool   `validate:"required"`
	ProductCategoryId string `validate:"required"`
	UnitName          string `validate:"required"`
	Description       string
}

type service struct {
	adminService admin.Service
	repository   Repository
}

func NewService(adminService admin.Service, repository Repository) Service {
	return &service{adminService, repository}
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
	admin, err := s.adminService.FindAdminByID(product.AdminId)
	if err != nil {
		return business.ErrNotHavePermission
	}

	data := NewProduct(
		product.Code,
		product.Name,
		product.Price,
		product.IsActive,
		product.ProductCategoryId,
		product.UnitName,
		product.Description,
		admin.ID,
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
		return business.ErrNotFound
	}
	err = validator.GetValidator().Struct(product)
	if err != nil {
		return business.ErrInvalidSpec
	}
	admin, err := s.adminService.FindAdminByID(updateProduct.AdminId)
	if err != nil {
		return business.ErrNotHavePermission
	}
	modifiedproduct := product.ModifyProduct(
		updateProduct.Name,
		updateProduct.Price,
		updateProduct.IsActive,
		updateProduct.ProductCategoryId,
		updateProduct.UnitName,
		updateProduct.Description,
		admin.ID,
		time.Now())

	return s.repository.UpdateProduct(modifiedproduct)
}

func (s *service) DeleteProduct(id string, adminId string) error {
	product, err := s.repository.FindProductById(id)
	if err != nil {
		return err
	} else if product == nil {
		return business.ErrNotFound
	} else if product.DeletedBy != "" {
		return business.ErrNotFound
	}
	admin, err := s.adminService.FindAdminByID(adminId)
	if err != nil {
		return business.ErrNotHavePermission
	}
	deleteProduct := product.DeleteProduct(
		time.Now(),
		admin.ID,
	)
	return s.repository.DeleteProduct(deleteProduct)
}
