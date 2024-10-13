package database

import (
	"github.com/EWRK0303/gin-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var gdb *gorm.DB

// 创建数据库的连接
func Setup() {
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})
	gdb = db
}

func GetDB() *gorm.DB {
	if gdb == nil {
		panic("database not setup")
	}
	return gdb
}
