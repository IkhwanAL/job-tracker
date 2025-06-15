package cryptography

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte, cost int) ([]byte, error) {

	bytes, err := bcrypt.GenerateFromPassword(password, cost)

	if err != nil {
		return nil, fmt.Errorf("failed to hash password %s", err)
	}

	return bytes, nil
}

func Compare(plainText []byte, hashPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPassword, plainText)
	return err
}
