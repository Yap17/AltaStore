package request

import (
	"AltaStore/business/category"
)

type InsertCategoryRequest struct {
	AdminId string `json:"adminid"`
	Code    string `json:"code"`
	Name    string `json:"name"`
}

func (i *InsertCategoryRequest) ToCategorySpec() *category.CategorySpec {
	var spec category.CategorySpec

	spec.AdminId = i.AdminId
	spec.Code = i.Code
	spec.Name = i.Name

	return &spec
}
