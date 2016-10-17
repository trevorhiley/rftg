package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/trevorhiley/rftg/models"
)

//CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := models.NewUser{}
	decoder.Decode(&user)
	_, err := user.CreateUser()
	if err != nil {
		errorString := err.Error()
		w.Write([]byte(errorString))
	} else {
		w.Write([]byte("success"))
	}
}
