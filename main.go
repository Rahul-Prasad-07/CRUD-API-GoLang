package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r) // get params
	// loop through movies and find id
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r) // get params
	// loop through movies and find id
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000)) // mock id - not safe
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	// set json content type
	// params
	// loop throgh the movies, rang
	// delete the movie with id that you have sent
	// add new movie - the movie that we send in the postman

	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r) // get params
	// loop through movies and find id
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

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
