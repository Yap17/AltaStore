package response

import (
	"AltaStore/business/admin"
)

//GetAdminResponse Get Admin by ID response payload
type GetAdminResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"name"`
	LastName  string `json:"Adminname"`
}

//NewGetAdminResponse construct GetAdminResponse
func NewGetAdminResponse(Admin admin.Admin) *GetAdminResponse {
	var getAdminResponse GetAdminResponse

	getAdminResponse.ID = Admin.ID
	getAdminResponse.Email = Admin.Email
	getAdminResponse.FirstName = Admin.FirstName
	getAdminResponse.LastName = Admin.LastName

	return &getAdminResponse
}
