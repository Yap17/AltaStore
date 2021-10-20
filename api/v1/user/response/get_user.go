package response

import (
	"AltaStore/business/user"
)

//GetUserResponse Get user by ID response payload
type GetUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"name"`
	LastName  string `json:"username"`
	HandPhone string `json:"handphone"`
	Address   string `json:"address"`
}

//NewGetUserResponse construct GetUserResponse
func NewGetUserResponse(user user.User) *GetUserResponse {
	var getUserResponse GetUserResponse

	getUserResponse.ID = user.ID
	getUserResponse.Email = user.Email
	getUserResponse.FirstName = user.FirstName
	getUserResponse.LastName = user.LastName
	getUserResponse.HandPhone = user.HandPhone
	getUserResponse.Address = user.Address

	return &getUserResponse
}
