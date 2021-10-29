package api

import (
	"AltaStore/api/middleware"
	"AltaStore/api/v1/admin"
	"AltaStore/api/v1/adminauth"
	"AltaStore/api/v1/category"
	"AltaStore/api/v1/checkout"
	"AltaStore/api/v1/checkoutpayment"
	"AltaStore/api/v1/product"
	"AltaStore/api/v1/purchasereceiving"
	"AltaStore/api/v1/shopping"
	"AltaStore/api/v1/user"
	"AltaStore/api/v1/userauth"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo,
	category *category.Controller,
	userController *user.Controller,
	adminController *admin.Controller,
	userAuthController *userauth.Controller,
	adminAuthController *adminauth.Controller,
	productController *product.Controller,
	shopping *shopping.Controller,
	checkout *checkout.Controller,
	purchaseController *purchasereceiving.Controller,
	paymentController *checkoutpayment.Controller,
) {
	if category == nil ||
		userController == nil ||
		userAuthController == nil ||
		adminAuthController == nil ||
		productController == nil ||
		shopping == nil ||
		purchaseController == nil ||
		paymentController == nil {
		panic("Invalid parameter")
	}

	// Add logger
	e.Use(middleware.MiddlewareLogger)

	// Custome response
	e.HTTPErrorHandler = func(e error, c echo.Context) {
		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		var response Response
		response.Code = http.StatusInternalServerError // defaul 500
		response.Message = "Internal Server Error"

		if he, ok := e.(*echo.HTTPError); ok {
			response.Code = he.Code
			response.Message = http.StatusText(he.Code)
		}

		c.Logger().Error(e)

		_ = c.JSON(response.Code, response)
	}

	// Routing
	regis := e.Group("v1/register")
	regis.POST("", userController.InsertUser)
	regis.POST("/admin", adminController.InsertAdmin)

	login := e.Group("v1/login")
	login.POST("", userAuthController.UserLogin)
	login.POST("/admin", adminAuthController.AdminLogin)

	admin := e.Group("v1/admins")
	admin.Use(middleware.JWTMiddleware())
	admin.PUT("/:id", adminController.UpdateAdmin)
	admin.DELETE("/:id", adminController.DeleteAdmin)
	admin.GET("/:id", adminController.FindAdminByID)
	admin.PUT("/:id/password", adminController.UpdateAdminPassword)

	cat := e.Group("v1/categories")
	cat.Use(middleware.JWTMiddleware())
	cat.GET("", category.GetAllCategory)
	cat.GET("/:id", category.FindCategoryById)
	cat.POST("", category.InsertCategory)
	cat.PUT("/:id", category.UpdateCategory)
	cat.DELETE("/:id", category.DeleteCategory)

	product := e.Group("v1/products")
	product.Use(middleware.JWTMiddleware())
	product.GET("", productController.GetAllProduct)
	//product.GET("/:id", productController.FindProductById)
	product.POST("", productController.InsertProduct)
	product.PUT("/:id", productController.UpdateProduct)
	product.DELETE("/:id", productController.DeleteProduct)

	// Routing shoping
	//e.GET("/v1/users/:id/shoppingcart", shopping.GetShoppingCartByUserId)

	shopCart := e.Group("/v1/shoppingcarts")
	shopCart.Use(middleware.JWTMiddleware())
	shopCart.POST("", shopping.NewShoppingCart)
	shopCart.GET("/:id", shopping.GetShopCartDetailById)
	shopCart.POST("/:id/products", shopping.NewItemInShopCart)
	shopCart.PUT("/:id/products/:productid", shopping.ModifyItemInShopCart)
	shopCart.DELETE("/:id/products/:productid", shopping.DeleteItemInShopCart)

	user := e.Group("v1/users")
	user.Use(middleware.JWTMiddleware())
	user.PUT("/:id", userController.UpdateUser)
	user.DELETE("/:id", userController.DeleteUser)
	user.GET("/:id", userController.FindUserByID)
	user.PUT("/:id/password", userController.UpdateUserPassword)
	user.GET("/:id/shoppingcart", shopping.GetShoppingCartByUserId)

	// Checkout
	c_out := e.Group("/v1/checkouts")
	c_out.POST("", checkout.NewCheckoutShoppingCart)
	c_out.GET("", checkout.GetAllCheckout)
	c_out.GET("/:id", checkout.GetCheckoutById)

	purchRec := e.Group("/v1/purchasereceivings")
	purchRec.Use(middleware.JWTMiddleware())
	purchRec.POST("", purchaseController.InsertPurchaseReceiving)
	purchRec.GET("", purchaseController.GetAllPurchaseReceiving)
	purchRec.GET("/:id", purchaseController.FindPurchaseReceivingById)
	purchRec.PUT("/:id", purchaseController.UpdatePurchaseReceiving)
	purchRec.DELETE("/:id", purchaseController.DeletePurchaseReceiving)

	payment := e.Group("/v1/payments")
	purchRec.Use(middleware.JWTMiddleware())
	payment.POST("", paymentController.Call)

	paymentCallback := e.Group("/v1/payments/notif")
	paymentCallback.GET("", paymentController.InsertPayment)

}
