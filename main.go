package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// main function
func main() {
	// add people
	addPeople()

	// create router
	router := mux.NewRouter()

	// register endpoints
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePeople).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePeople).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func addPeople() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Taro", Lastname: "Tanaka"})
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for _, person := range people {
		if person.ID == vars["id"] {
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePeople(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var person Person
	person.ID = vars["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func DeletePeople(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for i, person := range people {
		if person.ID == vars["id"] {
			people = append(people[:i], people[i+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
