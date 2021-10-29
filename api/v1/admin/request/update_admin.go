package request

import (
	"AltaStore/business/admin"
)

//UpdateAdminRequest create Admin request payload
type UpdateAdminRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//ToUpsertAdminSpec convert into Admin.UpsertAdminSpec object
func (req *UpdateAdminRequest) ToUpsertAdminSpec() *admin.UpdateAdminSpec {

	var updateAdminSpec admin.UpdateAdminSpec

	updateAdminSpec.FirstName = req.FirstName
	updateAdminSpec.LastName = req.LastName

	return &updateAdminSpec
}
