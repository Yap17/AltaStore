package api

import (
	"AltaStore/api/v1/auth"
	"AltaStore/api/v1/category"
	"AltaStore/api/v1/shopping"
	"AltaStore/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo,
	category *category.Controller,
	userController *user.Controller,
	authController *auth.Controller,
	shopping *shopping.Controller,
) {
	if category == nil || userController == nil {
		panic("Invalid parameter")
	}

	cat := e.Group("v1/categories")
	//cat.Use(middleware.JWTMiddleware())
	cat.GET("", category.GetAllCategory)
	cat.GET("/:id", category.FindCategoryById)
	cat.POST("", category.InsertCategory)
	cat.PUT("/:id", category.UpdateCategory)
	cat.DELETE("/:id", category.DeleteCategory)

	user := e.Group("v1/users")
	//userV1.GET("/:id", userController.FindUserByID)
	user.POST("", userController.InsertUser)
	user.PUT("/:id", userController.UpdateUser)
	user.DELETE("/:id", userController.DeleteUser)
	user.GET("/:id", userController.FindUserByID)
	user.PUT("/:id/password", userController.UpdateUserPassword)

	auth := e.Group("v1/users/login")
	auth.POST("", authController.Login)

	// Routing shoping
	e.GET("/v1/users/:id/shoppingcart", shopping.GetShoppingCartByUserId)

}
