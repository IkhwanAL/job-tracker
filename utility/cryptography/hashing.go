package cryptography

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) ([]byte, error) {

	bytes, err := bcrypt.GenerateFromPassword(password, 16)

	if err != nil {
		return nil, fmt.Errorf("failed to hash password %s", err)
	}

	return bytes, nil
}

func Compare(plainText []byte, hashPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPassword, plainText)

	return err
}
