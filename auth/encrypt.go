package auth

import "golang.org/x/crypto/bcrypt"

//EncryptPassword encrypts passwords
func EncryptPassword(password string) (encryptedPassword []byte, err error) {
	passwordBytes := []byte(password)

	encryptedPassword, error := bcrypt.GenerateFromPassword(passwordBytes, 10)

	if error != nil {
		return nil, error
	}
	return encryptedPassword, nil
}

//ValidatePassword validates passwords
func ValidatePassword(hashedPassword []byte, password string) (err error) {
	passwordBytes := []byte(password)

	err = bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)

	if err != nil {
		return err
	}

	return nil
}
