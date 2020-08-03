package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("./static/index.html")
	w.Write(file)
}

func Register(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("./static/register.html")
	w.Write(file)
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/register", Register).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
