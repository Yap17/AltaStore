package request

import (
	"AltaStore/business/category"
)

type InsertCategoryRequest struct {
	UserId string `json:"userid"`
	Code   string `json:"code"`
	Name   string `json:"name"`
}

func (i *InsertCategoryRequest) ToCategorySpec() *category.CategorySpec {
	var spec category.CategorySpec

	spec.UserId = i.UserId
	spec.Code = i.Code
	spec.Name = i.Name

	return &spec
}
