package request

import "AltaStore/business/category"

type UpdateCategoryRequest struct {
	AdminId string `json:"adminid"`
	Code    string `json:"code"`
	Name    string `json:"name"`
}

func (u *UpdateCategoryRequest) ToCategory() *category.CategorySpec {
	var spec category.CategorySpec

	spec.AdminId = u.AdminId
	spec.Code = u.Code
	spec.Name = u.Name

	return &spec
}
