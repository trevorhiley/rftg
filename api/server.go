package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/trevorhiley/rftg/datasvc"
	"github.com/trevorhiley/rftg/game"
)

func handler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(game.DiceTypeMap())
}

func redishandler(w http.ResponseWriter, r *http.Request) {
	newjson, _ := json.Marshal(game.DiceSideMap())
	jsonstring := string(newjson)
	returnedkey, _ := datasvc.Setkey(jsonstring)
	var mapdata map[int]string
	json.Unmarshal([]byte(returnedkey), &mapdata)
	fmt.Fprint(w, mapdata[1])
}

//Start starts the server
func Start() {
	datasvc.CreateClient()
	port := os.Getenv("PORT")
	http.HandleFunc("/redis", redishandler)
	http.HandleFunc("/", handler)

	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server starting on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
