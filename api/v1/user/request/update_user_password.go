package request

//UpdateUserRequest create User request payload
type UpdateUserPasswordRequest struct {
	password string `json:"password"`
}
