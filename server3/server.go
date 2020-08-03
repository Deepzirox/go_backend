package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHtml(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./static/index.html")
	if err != nil {
		panic(err)
	}
	w.Write(file)
}

func indexPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("lol"))
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	r.HandleFunc("/", indexHtml).Methods("GET")
	r.HandleFunc("/", indexPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}
