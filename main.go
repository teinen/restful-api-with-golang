package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// main function
func main() {
	// create router
	router := mux.NewRouter()

	// register endpoints
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePeople).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePeople).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request)    {}
func GetPerson(w http.ResponseWriter, r *http.Request)    {}
func CreatePeople(w http.ResponseWriter, r *http.Request) {}
func DeletePeople(w http.ResponseWriter, r *http.Request) {}
