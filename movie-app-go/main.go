package main

import (
	"encoding/json"
	f "fmt"
	"github.com/gorilla/mux"
	"log"
	rnd "math/rand"
	net "net/http"
	str "strconv"
)

type Director struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Year     string    `json:"year"`
	Director *Director `json:"director"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)

	movies = append(movies, Movie{
		ID:    "1",
		ISBN:  "123456789",
		Title: "A New Book",
		Year:  "2022",
		Director: &Director{
			ID:   "1",
			Name: "Talanted Director",
			Age:  35,
		},
	})

	router.HandleFunc("/movies", getAllMovies).Methods(net.MethodGet)
	router.HandleFunc("/movies/{id}", getMovieById).Methods(net.MethodGet)
	router.HandleFunc("/movies", saveMovie).Methods(net.MethodPost)
	router.HandleFunc("/movies/{id}", updateMovieById).Methods(net.MethodPut)
	router.HandleFunc("/movies/{id}", deleteMovieById).Methods(net.MethodDelete)

	f.Printf("Starting server at post 8000\n")
	log.Fatal(net.ListenAndServe(":8000", router))
}

func commonMiddleware(next net.Handler) net.Handler {
	return net.HandlerFunc(
		func(writer net.ResponseWriter, request *net.Request) {
			writer.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(writer, request)
		})
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getAllMovies(writer net.ResponseWriter, request *net.Request) {
	err := json.NewEncoder(writer).Encode(movies)
	handleError(err)
}

func deleteMovieById(writer net.ResponseWriter, request *net.Request) {
	params := mux.Vars(request)

	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}

	err := json.NewEncoder(writer).Encode(movies)
	handleError(err)
}

func getMovieById(writer net.ResponseWriter, request *net.Request) {
	params := mux.Vars(request)

	for _, item := range movies {
		if item.ID == params["id"] {
			err := json.NewEncoder(writer).Encode(item)
			handleError(err)
			return
		}
	}
}

func saveMovie(writer net.ResponseWriter, request *net.Request) {
	var newMovie Movie
	err := json.NewDecoder(request.Body).Decode(&newMovie)
	handleError(err)

	newMovie.ID = str.Itoa(rnd.Intn(1_000_000))
	movies = append(movies, newMovie)

	err = json.NewEncoder(writer).Encode(movies)
	handleError(err)
}

func updateMovieById(writer net.ResponseWriter, request *net.Request) {
	params := mux.Vars(request)
	id := params["id"]

	for i, item := range movies {
		if item.ID == id {
			movies = append(movies[:i], movies[i+1:]...)

			var updatedMovie Movie
			err := json.NewDecoder(request.Body).Decode(&updatedMovie)
			handleError(err)

			updatedMovie.ID = id
			movies = append(movies, updatedMovie)

			err = json.NewEncoder(writer).Encode(movies)
			handleError(err)
			return
		}
	}
}
