package response

import "AltaStore/business/shopping"

func GetOneResponse(data *shopping.ShoppCart) *ShoppData {
	return &ShoppData{
		ID:          data.ID,
		IsCheckOut:  data.IsCheckOut,
		Description: data.Description,
		UpdatedAt:   data.UpdatedAt,
	}
}
