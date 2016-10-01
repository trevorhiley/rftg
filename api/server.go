package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/trevorhiley/rftg/game"
)

func handler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(game.DiceTypeMap())
}

//Start starts the server
func Start() {
	fmt.Print("Server starting")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
