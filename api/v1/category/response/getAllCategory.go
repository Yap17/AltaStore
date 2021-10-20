package response

import (
	"AltaStore/business/category"
	"time"
)

type GetCategoryResponse struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewGetCategoryResponse(category category.Category) *GetCategoryResponse {
	var res GetCategoryResponse

	res.ID = category.ID
	res.Code = category.Code
	res.Name = category.Name
	res.UpdatedAt = category.UpdatedAt

	return &res
}
