package request

import (
	"AltaStore/business/user"
)

//InsertUserRequest create User request payload
type InsertUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

//ToUpsertUserSpec convert into User.UpsertUserSpec object
func (req *InsertUserRequest) ToUpsertUserSpec() *user.InsertUserSpec {
	var insertUserSpec user.InsertUserSpec

	insertUserSpec.Email = req.Email
	insertUserSpec.FirstName = req.FirstName
	insertUserSpec.LastName = req.LastName
	insertUserSpec.Password = req.Password

	return &insertUserSpec
}
