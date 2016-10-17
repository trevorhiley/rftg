package models

import (
	"testing"

	"github.com/trevorhiley/rftg/db"
)

func TestCreateUser(t *testing.T) {
	newUser := NewUser{
		UserName:  "billandted",
		Password:  "bobjoe",
		FirstName: "Bill",
		LastName:  "Ted",
		Email:     "bill.ted@gmail.com",
	}

	if newUser.Email != "bill.ted@gmail.com" {
		t.Error("email not set")
	}

	db.CreateClient()

	_, err := newUser.CreateUser()

	if err != nil {

		t.Error("error creating user")
	}

	user, _ := GetUser("billandted")

	err = user.Authenticate("bobjoe")

	if err != nil {
		t.Error("password auth not working")
	}

	err = user.Authenticate("tdob$")

	if err == nil {
		t.Error("wrong password still working")
	}

}
