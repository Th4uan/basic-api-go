package database

import "gorm.io/gorm"

type User struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}
