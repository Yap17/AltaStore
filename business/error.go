package business

import "errors"

var (
	//ErrInternalServerError Error caused by system error
	ErrInternalServerError = errors.New("Internal Server Error")

	//ErrHasBeenModified Error when update item that has been modified
	ErrHasBeenModified = errors.New("Data has been modified")

	//ErrNotFound Error when item is not found
	ErrNotFound = errors.New("Data was not found")

	//ErrOldPasswordMisMatch Error when old password wrong
	ErrPasswordMisMatch = errors.New("Wrong Password")

	//ErrUserDeleted Error when user deleted
	ErrUserDeleted = errors.New("User Deleted")

	//ErrInvalidSpec Error when data given is not valid on update or insert
	ErrInvalidSpec = errors.New("Given spec is not valid")

	//ErrLoginFailed Error when Login Failed
	ErrLoginFailed = errors.New("Login Failed")
)
