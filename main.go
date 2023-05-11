package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"math/rand"
	"encoding/json"

	"github.com/gorilla/mux"
)

// for routing : for beginner phase we are not using complex concept like routing

// we are not using databse, we are using struct for store data
// struct of : movie and Director

type Movie struct {
	ID       string    `json:"id`        // json for encode data
	Isbn     string    `json:"isbn"`     // Isbn is uniuqe id for every movie
	Title    string    `json:"title"`    // title of movie
	Director *Director `json:"director"` // director of movie : * used for other struct is associated with that

}

// movie and director are going to associated with each other : every movie has director
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {

	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Star Wars-1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "454555", Title: "Star Wars-2", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
