package request

import (
	"AltaStore/business/admin"
)

//UpdateAdminRequest create Admin request payload
type UpdateAdminPasswordRequest struct {
	NewPassword string `json:"newpassword"`
	OldPassword string `json:"oldpassword"`
}

//ToUpsertAdminSpec convert into Admin.UpsertAdminSpec object
func (req *UpdateAdminPasswordRequest) ToUpsertAdminSpec() *admin.UpdateAdminPasswordSpec {

	var updateAdminPasswordSpec admin.UpdateAdminPasswordSpec

	updateAdminPasswordSpec.NewPassword = req.NewPassword
	updateAdminPasswordSpec.OldPassword = req.OldPassword

	return &updateAdminPasswordSpec
}
