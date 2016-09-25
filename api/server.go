package main

import (
	"encoding/json"
	"net/http"
)

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	jsonResp := Response2{Page: 1, Fruits: []string{"apple", "banana"}}

	json.NewEncoder(w).Encode(jsonResp)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
