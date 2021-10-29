package request

import "AltaStore/business/category"

type UpdateCategoryRequest struct {
	Name string `json:"name"`
}

func (u *UpdateCategoryRequest) ToCategory() *category.CategorySpec {
	var spec category.CategorySpec
	spec.Name = u.Name

	return &spec
}
