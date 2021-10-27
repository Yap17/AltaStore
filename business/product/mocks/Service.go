// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	product "AltaStore/business/product"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// DeleteProduct provides a mock function with given fields: id, adminId
func (_m *Service) DeleteProduct(id string, adminId string) error {
	ret := _m.Called(id, adminId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(id, adminId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindProductById provides a mock function with given fields: id
func (_m *Service) FindProductById(id string) (*product.Product, error) {
	ret := _m.Called(id)

	var r0 *product.Product
	if rf, ok := ret.Get(0).(func(string) *product.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllProduct provides a mock function with given fields:
func (_m *Service) GetAllProduct() (*[]product.Product, error) {
	ret := _m.Called()

	var r0 *[]product.Product
	if rf, ok := ret.Get(0).(func() *[]product.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllProductByParameter provides a mock function with given fields: id, isActive, categoryName, code, name
func (_m *Service) GetAllProductByParameter(id string, isActive string, categoryName string, code string, name string) (*[]product.Product, error) {
	ret := _m.Called(id, isActive, categoryName, code, name)

	var r0 *[]product.Product
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) *[]product.Product); ok {
		r0 = rf(id, isActive, categoryName, code, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, string) error); ok {
		r1 = rf(id, isActive, categoryName, code, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertProduct provides a mock function with given fields: Product
func (_m *Service) InsertProduct(Product *product.InsertProductSpec) error {
	ret := _m.Called(Product)

	var r0 error
	if rf, ok := ret.Get(0).(func(*product.InsertProductSpec) error); ok {
		r0 = rf(Product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProduct provides a mock function with given fields: id, Product
func (_m *Service) UpdateProduct(id string, Product *product.UpdateProductSpec) error {
	ret := _m.Called(id, Product)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *product.UpdateProductSpec) error); ok {
		r0 = rf(id, Product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}