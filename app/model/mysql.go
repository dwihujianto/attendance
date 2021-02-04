package model

import (
	"fmt"

	"attendance/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit(config *config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&Presence{})

	DB = db
}
