package auth

//User struct
type User struct {
	UserName       string
	HashedPassword []byte
	FirstName      string
	LastName       string
	Email          string
}

//ValidateUserPassword validates users password
func (u *User) ValidateUserPassword(password string) (err error) {
	err = ValidatePassword(u.HashedPassword, password)
	if err != nil {
		return err
	}
	return nil
}

//HashPassword hash password and add it to user object
func (u *User) HashPassword(password string) (err error) {
	hashedPassword, err := EncryptPassword(password)

	if err != nil {
		return err
	}

	u.HashedPassword = hashedPassword

	return

}
