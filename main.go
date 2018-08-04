package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/** Person struct */
type Person struct {
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
}

/** Address struct */
type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

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
