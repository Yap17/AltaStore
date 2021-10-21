package migration

import (
	"AltaStore/modules/admin"
	"AltaStore/modules/category"
	"AltaStore/modules/product"
	"AltaStore/modules/user"

	"gorm.io/gorm"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&category.ProductCategory{},
		&user.User{},
		&product.Product{},
		&admin.Admin{},
	)
}
