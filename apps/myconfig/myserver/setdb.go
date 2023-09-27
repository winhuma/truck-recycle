package myserver

import (
	"truck-unii-app/apps/myconfig/mymodels"

	"gorm.io/gorm"
)

func InitDB(db *gorm.DB) {
	db.AutoMigrate(mymodels.DBTruckProfile{})
}
