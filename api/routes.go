package api

import (
	"AltaStore/api/v1/category"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, category *category.Controller) {
	if category == nil {
		panic("Invalid parameter")
	}

	cat := e.Group("v1/categories")
	cat.GET("", category.GetAllCategory)
}
