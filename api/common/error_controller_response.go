package common

import "net/http"

const (
	BadRequestCode responseCode = "400"
	ForbiddenCode  responseCode = "403"
	NotFoundCode   responseCode = "404"
)

func BadRequestResponse() (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{BadRequestCode, "Bad Request", map[string]interface{}{}}
}

func ForbiddenResponse() (int, ControllerResponse) {
	return http.StatusForbidden, ControllerResponse{ForbiddenCode, "Forbidden", map[string]interface{}{}}
}

func NotFoundResponse() (int, ControllerResponse) {
	return http.StatusNotFound, ControllerResponse{NotFoundCode, "Not Found", map[string]interface{}{}}
}
