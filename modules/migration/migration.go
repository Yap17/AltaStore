package migration

import (
	"AltaStore/modules/category"
	"AltaStore/modules/user"

	"gorm.io/gorm"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&category.ProductCategory{}, &user.User{})
}
