package request

import (
	"AltaStore/business/admin"
)

//InsertAdminRequest create Admin request payload
type InsertAdminRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

//ToUpsertAdminSpec convert into Admin.UpsertAdminSpec object
func (req *InsertAdminRequest) ToUpsertAdminSpec() *admin.InsertAdminSpec {
	var insertAdminSpec admin.InsertAdminSpec

	insertAdminSpec.Email = req.Email
	insertAdminSpec.FirstName = req.FirstName
	insertAdminSpec.LastName = req.LastName
	insertAdminSpec.Password = req.Password

	return &insertAdminSpec
}
