package migration

import (
	"AltaStore/modules/admin"
	"AltaStore/modules/category"
	"AltaStore/modules/checkout"
	"AltaStore/modules/product"
	"AltaStore/modules/shopping"
	"AltaStore/modules/shoppingdetail"
	"AltaStore/modules/user"

	"gorm.io/gorm"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&category.ProductCategory{},
		&shopping.ShoppingCart{},
		&user.User{},
		&product.Product{},
		&admin.Admin{},
		&shoppingdetail.ShoppingCartDetail{},
		&checkout.Checkout{},
	)
}
