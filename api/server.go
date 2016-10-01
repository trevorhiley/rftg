package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/trevorhiley/rftg/game"
)

func handler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(game.DiceTypeMap())
}

//Start starts the server
func Start() {
	port := os.Getenv("PORT")
	fmt.Print("Server starting")
	http.HandleFunc("/", handler)
	if port != "" {
		http.ListenAndServe(":"+port, nil)
	} else {
		http.ListenAndServe(":8080", nil)
	}
}
