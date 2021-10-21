package api

import (
	"AltaStore/api/v1/admin"
	"AltaStore/api/v1/adminauth"
	"AltaStore/api/v1/category"
	"AltaStore/api/v1/product"
	"AltaStore/api/v1/user"
	"AltaStore/api/v1/userauth"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo,
	category *category.Controller,
	userController *user.Controller,
	adminController *admin.Controller,
	userAuthController *userauth.Controller,
	adminAuthController *adminauth.Controller,
	productController *product.Controller,
) {
	if category == nil ||
		userController == nil ||
		userAuthController == nil ||
		adminAuthController == nil ||
		productController == nil {
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
	user.POST("", userController.InsertUser)
	user.PUT("/:id", userController.UpdateUser)
	user.DELETE("/:id", userController.DeleteUser)
	user.GET("/:id", userController.FindUserByID)
	user.PUT("/:id/password", userController.UpdateUserPassword)

	authUser := e.Group("v1/users/login")
	authUser.POST("", userAuthController.UserLogin)

	admin := e.Group("v1/admins")
	admin.POST("", adminController.InsertAdmin)
	admin.PUT("/:id", adminController.UpdateAdmin)
	admin.DELETE("/:id", adminController.DeleteAdmin)
	admin.GET("/:id", adminController.FindAdminByID)
	admin.PUT("/:id/password", adminController.UpdateAdminPassword)

	authAdmin := e.Group("v1/admins/login")
	authAdmin.POST("", adminAuthController.AdminLogin)

	product := e.Group("v1/products")
	//product.Use(middleware.JWTMiddleware())
	product.GET("", productController.GetAllProduct)
	//product.GET("/:id", productController.FindProductById)
	product.POST("", productController.InsertProduct)
	product.PUT("/:id", productController.UpdateProduct)
	product.DELETE("/:id", productController.DeleteProduct)

}
