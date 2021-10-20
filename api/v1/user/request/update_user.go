package request

import (
	"AltaStore/business/user"
)

//UpdateUserRequest create User request payload
type UpdateUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	HandPhone string `json:"handphone"`
	Address   string `json:"address"`
}

//ToUpsertUserSpec convert into User.UpsertUserSpec object
func (req *UpdateUserRequest) ToUpsertUserSpec() *user.UpdateUserSpec {

	var updateUserSpec user.UpdateUserSpec

	updateUserSpec.Email = req.Email
	updateUserSpec.FirstName = req.FirstName
	updateUserSpec.LastName = req.LastName
	updateUserSpec.HandPhone = req.HandPhone
	updateUserSpec.Address = req.Address

	return &updateUserSpec
}
