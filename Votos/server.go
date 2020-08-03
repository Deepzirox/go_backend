package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	engine "./models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// statics
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// routes
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/", indexAddEmail).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))

}

func indexAddEmail(w http.ResponseWriter, r *http.Request) {

	Nemail := r.FormValue("emailform")
	x1, x2 := engine.ValidateEmail(Nemail)
	if !(x1 || x2) {
		w.Write([]byte("none"))
		return
	}
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	newJSON := engine.User{
		Email: Nemail,
		IP:    forwarded,
	}
	err := engine.SaveUser(newJSON)
	if err != nil {
		w.Write([]byte("Invalid email fuck**ing hacker :)"))
		return
	}
	json.NewEncoder(w).Encode(newJSON)
}

func index(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("./static/index.html")
	w.Write(file)
}
