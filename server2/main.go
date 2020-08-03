package main

import (
	"encoding/json"
	"log"
	"net/http"

	KEY "./keyGenerator"
	"github.com/gorilla/mux"
)

type Player struct {
	ID      string   `json:"id,omitempty"`
	Level   string   `json:"level,omitempty"`
	Trofeos *Trofeos `json:"trofeos,omitempty"`
}

type Trofeos struct {
	Competitivos string `json:"competitivos,omitempty"`
	Personales   string `json:"personales,omitempty"`
}

var players []Player

func helloworld(w http.ResponseWriter, res *http.Request) {
	player1 := Player{
		ID:    KEY.KeyGen([]uint{2, 1, 1, 2}),
		Level: "30",
		Trofeos: &Trofeos{
			Competitivos: "20",
			Personales:   "5",
		},
	}

	players = append(players, player1)
	json.NewEncoder(w).Encode(players)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/jugadores", helloworld).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Fatal(http.ListenAndServe(":3000", router))
}
