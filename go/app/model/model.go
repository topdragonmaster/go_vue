package model

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"AUTO_INCREMENT,primary_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"userName"`
	UserRole string `gorm:"type:ENUM('0', '1', '2', '3');default:'0'"`
}

type Category struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Title string `json:"title"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Category{}, &User{})
	return db
}
