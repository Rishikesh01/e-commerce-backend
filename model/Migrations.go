package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

//singleton
func Init() *gorm.DB {

	if db != nil {
		return db
	}
	dsn := "postgres://boris@localhost:5432/gorm"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("error:", err)
	}
	if err = db.AutoMigrate(&User{},
		&Product{},
		&UserOrderHistory{},
		&Seller{},
		&Address{}); err != nil {
		log.Println("error:", err)
	}
	return db
}
