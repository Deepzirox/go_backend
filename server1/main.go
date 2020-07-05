package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person - struct that contains person info
/*
	ID: id of user
	firstname: firstname of user
	lastname: lastname of user
	Address: pointer to struct Address that contains {city, state}
*/
type Person struct {
	ID        string   `json:"id,omitempPersonty"`
	FirstName string   `json:"firstname,omitempPersonty"`
	LastName  string   `json:"lastname,omitempPersonty"`
	Address   *Address `json:"address,omitempPersonty"`
}

type Address struct {
	City  string `json:"city,omitempPersonty"`
	State string `json:"state,omitempPersonty"`
}

var people []Person

func getPeopleEnd(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.Write([]byte("<h1>INVALIDOOOOO</h1>"))
}

func addPeople(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var tempPerson Person

	_ = json.NewDecoder(req.Body).Decode(&tempPerson)
	tempPerson.ID = params["id"]
	people = append(people, tempPerson)
	json.NewEncoder(w).Encode(people)
}

func deletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	newPeoples := make([]Person, len(people))
	for index, item := range people {
		if item.ID == params["id"] {
			newPeoples = people[0:index]
			newPeoples = append(newPeoples, people[index+1:]...)
			people = newPeoples
			json.NewEncoder(w).Encode(people)
			return
		}
	}
	w.Write([]byte("404 Not found"))
}

func routes() {
	// init router
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./statics")))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/people", getPeopleEnd).Methods("GET")
	router.HandleFunc("/people/{id}", getPerson).Methods("GET")
	router.HandleFunc("/people/{id}", addPeople).Methods("POST")
	router.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")

	routes()

	log.Fatal(http.ListenAndServe(":3000", router))

}
