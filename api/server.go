package main

import (
	"encoding/json"
	"net/http"

	"github.com/trevorhiley/rftg/game"
)

func handler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(game.DiceTypeMap())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
