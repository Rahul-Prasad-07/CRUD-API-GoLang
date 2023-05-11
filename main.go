package main

// for formatting and printing
// for logging
// for encode my data when send to postman
// to recate new movie at random
// for creating server

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
