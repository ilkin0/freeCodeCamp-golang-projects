package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ilkinmehdiyev.com/begginer-freecodecamp/book-mng-app-go/pkg/utils"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:password@tcp(localhost:3306)/go_book_mng?createDatabaseIfNotExist=true&charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	utils.HandleErrorPanic(err)
	db = d
}

func GetDB() *gorm.DB {
	return db
}
