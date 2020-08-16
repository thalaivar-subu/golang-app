package database

import (
	"github.com/jinzhu/gorm"
	"github.com/thalaivar-subu/golang-app/backend/structs"
)

// Migrate -> Migrates table
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&structs.User{})
}
