package migration

import (
	"AltaStore/modules/admin"
	"AltaStore/modules/category"
	"AltaStore/modules/product"
	"AltaStore/modules/purchasereceiving"

	"AltaStore/modules/purchasereceivingdetail"
	"AltaStore/modules/shopping"
	"AltaStore/modules/shoppingdetail"
	"AltaStore/modules/user"

	"gorm.io/gorm"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&category.ProductCategory{},
		&user.User{},
		&product.Product{},
		&admin.Admin{},
		&shopping.ShoppingCart{},
		&shoppingdetail.ShoppingCartDetail{},
		&purchasereceiving.PurchaseReceiving{},
		&purchasereceivingdetail.PurchaseReceivingDetail{},
	)
}
