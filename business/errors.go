package business

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrInternalServer = errors.New("Internal_Server_Error")

	ErrHasBeenModified = gorm.ErrInvalidData

	ErrNotFound = gorm.ErrRecordNotFound

	ErrInvalidSpec = errors.New("Given_Spec_Is_Not_Valid")

	ErrPasswordMisMatch = errors.New("Wrong Password")

	ErrUserDeleted = errors.New("User Deleted")

	ErrLoginFailed = errors.New("Login Failed")
)
