package migration

import (
	"AltaStore/modules/category"

	"gorm.io/gorm"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&category.ProductCategory{})
}
