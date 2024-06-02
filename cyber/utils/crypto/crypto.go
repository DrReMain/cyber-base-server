package cutils_crypto

import "golang.org/x/crypto/bcrypt"

func Crypt(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), 5)
	return string(hashed), err
}

func Verify(origin, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(origin))
	if err != nil {
		return false
	}
	return true
}
