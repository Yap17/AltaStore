package request

import (
	"AltaStore/business/admin"
)

//UpdateAdminRequest create Admin request payload
type UpdateAdminRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//ToUpsertAdminSpec convert into Admin.UpsertAdminSpec object
func (req *UpdateAdminRequest) ToUpsertAdminSpec() *admin.UpdateAdminSpec {

	var updateAdminSpec admin.UpdateAdminSpec

	updateAdminSpec.Email = req.Email
	updateAdminSpec.FirstName = req.FirstName
	updateAdminSpec.LastName = req.LastName

	return &updateAdminSpec
}
