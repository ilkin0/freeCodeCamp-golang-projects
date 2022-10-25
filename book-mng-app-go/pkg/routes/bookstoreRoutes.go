package routes

import (
	"github.com/gorilla/mux"
	"ilkinmehdiyev.com/begginer-freecodecamp/book-mng-app-go/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.GetAllBooks)
}
