package main

import (
	"github.com/gorilla/mux"
	"ilkinmehdiyev.com/begginer-freecodecamp/book-mng-app-go/pkg/routes"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9090", router))
}
