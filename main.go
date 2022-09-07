package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies,
		Movie{ID: "1", Isbn: "123", Title: "Agustos", Director: &Director{Name: "furkan", Surname: "Demirel"}})
	movies = append(movies,
		Movie{"2", "123", "Interseller", &Director{"Ahmed", "Sertik"}})
	movies = append(movies,
		Movie{"3", "5678", "Sonbahar", &Director{"Kemal", "Soyler"}})

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getOneMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("yeap its listening")
	log.Fatal(http.ListenAndServe(":8182", r))
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}

}
func getOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	err := json.NewEncoder(w).Encode(movie)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("added")

}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100))
			movies = append(movies, movie)
			err := json.NewEncoder(w).Encode(movies)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("updated")
		}
	}
}
