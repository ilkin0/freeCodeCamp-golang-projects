package models

import (
	"gorm.io/gorm"
	"ilkinmehdiyev.com/begginer-freecodecamp/book-mng-app-go/pkg/config"
	"ilkinmehdiyev.com/begginer-freecodecamp/book-mng-app-go/pkg/utils"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	ISBN        string `json:"isbn"`
	Year        int    `json:"year"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	err := db.AutoMigrate(&Book{})
	utils.HandleErrorPanic(err)
}
