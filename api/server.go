package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/trevorhiley/rftg/db"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server is up"))
}

//Start starts the server
func Start() {
	db.CreateClient()
	port := os.Getenv("PORT")
	r := getRoutes()
	http.Handle("/", r)

	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server starting on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
