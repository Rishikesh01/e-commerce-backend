package repository

import (
	"github.com/Rishikesh01/amazon-clone-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var _db *gorm.DB

func Init() *gorm.DB {

	if _db != nil {
		return _db
	}
	dsn := "postgres://boris@localhost:5432/ecom"
	var err error
	_db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("error:", err)
	}
	if err = _db.AutoMigrate(&model.User{},
		&model.Seller{},
		&model.Product{},
		&model.Billing{},
		&model.UserOrderHistory{},
		&model.ProductSeller{}); err != nil {
		log.Println("error:", err)
	}
	return _db
}
