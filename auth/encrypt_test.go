package auth

import (
	"testing"
)

func TestEncryptpassword(t *testing.T) {
	encryptedPassword, err := EncryptPassword("Trevor")

	if encryptedPassword == nil {
		t.Error("password didn't generate")
	}

	if err != nil {
		t.Error("password encryption through error and shouldn't have")
	}

	err = ValidatePassword(encryptedPassword, "Trevor")

	if err != nil {
		t.Error("Password matching not working with valid pass")
	}

	err = ValidatePassword(encryptedPassword, "Julie")

	if err == nil {
		t.Error("Password matching not working with invalid password", err)
	}
}
