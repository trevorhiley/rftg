package models

import (
	"encoding/json"
	"errors"

	"github.com/trevorhiley/rftg/auth"
	"github.com/trevorhiley/rftg/db"
)

//User struct
type User struct {
	UserName       string
	HashedPassword []byte
	FirstName      string
	LastName       string
	Email          string
}

//NewUser struct to temporarily hold new user def
type NewUser struct {
	UserName       string
	HashedPassword []byte
	FirstName      string
	LastName       string
	Email          string
	Password       string
}

//GetUser info
func GetUser(userName string) (user User, err error) {
	user, err = getUserFromDb(userName)
	if err != nil {
		return user, err
	}
	return user, nil
}

//CreateUser using NewUser
func (n *NewUser) CreateUser() (user User, err error) {
	if n.Password == "" {
		return user, errors.New("You must provide a password")
	}
	hashPassword, _ := hashPassword(n.Password)
	user = User{UserName: n.UserName,
		HashedPassword: hashPassword,
		FirstName:      n.FirstName,
		LastName:       n.LastName,
		Email:          n.Email,
	}

	err = createUserInDb(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

//Authenticate validates users password
func (u User) Authenticate(password string) (err error) {
	err = auth.ValidatePassword(u.HashedPassword, password)
	if err != nil {
		return err
	}
	return nil
}

func (n *NewUser) validateNewUser() (err []error) {
	return nil
}

//HashPassword hash password and add it to user object
func hashPassword(password string) (hashedPassword []byte, err error) {
	hashedPassword, err = auth.EncryptPassword(password)

	if err != nil {
		return hashedPassword, err
	}

	return hashedPassword, nil

}

//GetUser gets user from redis
func getUserFromDb(userName string) (user User, err error) {
	returnedkey, _ := db.Getkey("user-" + userName)
	if returnedkey == "" {
		return user, errors.New("username not found")
	}

	json.Unmarshal([]byte(returnedkey), &user)

	if user.UserName == "" {
		return user, errors.New("key was not in the form of a user")
	}

	return user, nil
}

//CreateUser creates user in db
func createUserInDb(inputUser User) (err error) {
	newjson, _ := json.Marshal(inputUser)
	jsonstring := string(newjson)
	_, err = db.Setkey("user-"+inputUser.UserName, jsonstring)
	if err != nil {
		return errors.New("User not found")
	}
	return
}
