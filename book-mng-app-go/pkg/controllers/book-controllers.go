package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"ilkinmehdiyev.com/begginer-freecodecamp/book-mng-app-go/pkg/models"
	"ilkinmehdiyev.com/begginer-freecodecamp/book-mng-app-go/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()

	res, err := json.Marshal(newBooks)
	utils.HandleErrorPanic(err)

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	utils.HandleErrorPanic(err)
}