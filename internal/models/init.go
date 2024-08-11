package models

import (
	// "gorm.io/driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:abcdi2124Jcke23@tcp(192.168.1.8:3306)/meeting?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&RoomBasic{}, &RoomUser{}, &UserBasic{})
	DB = db
}
