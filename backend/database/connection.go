package database

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/thalaivar-subu/golang-app/backend/config"

	// Gorm Blank Import
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ConnectMysql -> Connects Mysql
func ConnectMysql() *gorm.DB {
	db, err := gorm.Open("mysql", config.Config["mysql"])
	if err != nil {
		glog.Info("Connection Failed to Open", err)
	} else {
		glog.Info("Connection Established")
		Migrate(db)
	}
	return db
}
