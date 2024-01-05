package main
import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)
type Movie struct{
ID string `json:"id"`
Isbn string `json:"isbn"`
Title string `json:"title"`
Director *Director `json:"director"`
}
type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`

}

var movies []Movie
func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main(){
r:=mux.NewRouter()

movies=append(movies, Movie{ID: "1", Isbn: "23332", Title: "Avatar",Director:&Director{FirstName: "Deribewe", LastName: "Shimelis"}})
movies=append(movies, Movie{ID: "2", Isbn: "49355", Title: "Movie 2",Director:&Director{FirstName: "Daniel", LastName: "Shimelis"}})
movies=append(movies, Movie{ID: "3", Isbn: "83332", Title: "movi 3",Director:&Director{FirstName: "Endale", LastName: "Shimelis"}})
r.HandleFunc("/movies",getMovies).Methods("GET")
r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
r.HandleFunc("/createMovie",createMovie).Methods("POST")
r.HandleFunc("/update/{id}",updateMovie).Methods("PUT")
r.HandleFunc("/delete/{id}",deleteMovie).Methods("DELETE")
fmt.Println("Starting Server at port :8000")
log.Fatal(http.ListenAndServe(":8000", r))

}