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
	Year        int64  `json:"year"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	err := db.AutoMigrate(&Book{})
	utils.HandleErrorPanic(err)
}

func CreateBook(b *Book) *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.First(&getBook, Id)
	return &getBook, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Delete(&book, Id)
	return book
}
