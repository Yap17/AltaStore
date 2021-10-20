package request

import (
	"AltaStore/business/user"
)

//UpdateUserRequest create User request payload
type UpdateUserPasswordRequest struct {
	NewPassword string `json:"newpassword"`
	OldPassword string `json:"oldpassword"`
}

//ToUpsertUserSpec convert into User.UpsertUserSpec object
func (req *UpdateUserPasswordRequest) ToUpsertUserSpec() *user.UpdateUserPasswordSpec {

	var updateUserPasswordSpec user.UpdateUserPasswordSpec

	updateUserPasswordSpec.NewPassword = req.NewPassword
	updateUserPasswordSpec.OldPassword = req.OldPassword

	return &updateUserPasswordSpec
}
