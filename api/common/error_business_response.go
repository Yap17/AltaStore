package common

import (
	"AltaStore/business"
	"net/http"
)

const (
	errInternalServerError responseCode = "500"
	errNotFound            responseCode = "404"
	errHasBeenModified     responseCode = "400"
	errNotHavePermission   responseCode = "400"
	errPasswordMisMatch    responseCode = "400"
	// errInvalidSpec         responseCode = ""
)

// Mengembalikan respons status dari permintaan
func NewBusinessErrorResponse(err error) (int, ControllerResponse) {
	return errorMapping(err)
}

func errorMapping(err error) (int, ControllerResponse) {
	switch err {
	default:
		return newInternalServerErrorResponse()

	case business.ErrNotFound:
		return newNotFoundResponse()

	case business.ErrHasBeenModified:
		return newHasBeenModifiedResponse()

	case business.ErrNotHavePermission:
		return newNotHavePermission()

	case business.ErrPasswordMisMatch:
		return newErrPasswordMisMatch()
	}
}

func newInternalServerErrorResponse() (int, ControllerResponse) {
	return http.StatusInternalServerError,
		ControllerResponse{errInternalServerError, "Internal Server Error", map[string]interface{}{}}
}

func newNotFoundResponse() (int, ControllerResponse) {
	return http.StatusNotFound,
		ControllerResponse{errNotFound, "Data Not Found", map[string]interface{}{}}
}

func newHasBeenModifiedResponse() (int, ControllerResponse) {
	return http.StatusBadRequest,
		ControllerResponse{errHasBeenModified, "Data Has Been Modified", map[string]interface{}{}}
}

func newNotHavePermission() (int, ControllerResponse) {
	return http.StatusBadRequest,
		ControllerResponse{errHasBeenModified, "Not Have Permission", map[string]interface{}{}}
}

func newErrPasswordMisMatch() (int, ControllerResponse) {
	return http.StatusBadRequest,
		ControllerResponse{errHasBeenModified, "Wrong Password", map[string]interface{}{}}
}
