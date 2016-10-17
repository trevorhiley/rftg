package api

import (
	"github.com/gorilla/mux"
	"github.com/trevorhiley/rftg/api/controllers"
)

func getRoutes() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/users", controllers.CreateUser).Methods("Post")
	return r
}
