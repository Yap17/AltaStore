package api

import (
	"AltaStore/api/v1/category"
	"AltaStore/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, category *category.Controller, userController *user.Controller) {
	if category == nil || userController == nil {
		panic("Invalid parameter")
	}

	cat := e.Group("v1/categories")
	cat.GET("", category.GetAllCategory)

	user := e.Group("v1/users")
	//userV1.GET("/:id", userController.FindUserByID)
	user.POST("", userController.InsertUser)
	user.PUT("/:id", userController.UpdateUser)
	user.DELETE("/:id", userController.DeleteUser)
}
