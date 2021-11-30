package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// "Person type" (tipo um objeto)
type Person struct {
	ID        string `json:"id"`
	Firstname string `json:"sistema"`
	Lastname  string `json:"torre"`
	Address   bool   `json:"criticidade"`
}

var people []Person

// Setup
func setupHeader(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// GetPeople mostra todos os contatos da variável people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	setupHeader(w, r)
	json.NewEncoder(w).Encode(people)
}

// GetPerson mostra apenas um contato
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// CreatePerson cria um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	setupHeader(w, r)
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson deleta um contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

// função principal para executar a api
func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "NPS", Lastname: "APPIA", Address: true})
	people = append(people, Person{ID: "2", Firstname: "Worx", Lastname: "APPIA", Address: false})
	router.HandleFunc("/systens", GetPeople).Methods("GET", "OPTIONS")
	router.HandleFunc("/systens/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/systens", CreatePerson).Methods("POST", "OPTIONS")
	router.HandleFunc("/systens/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3001", router))
}
