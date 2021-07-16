package vos

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var ErrInvalidSecret = errors.New("invalid secret")

type Secret string

type Hash interface {
	HashAndSalt(secret Secret) string
}

func IsValidSecret(secret Secret) error {
	if secret == "" {
		return ErrInvalidSecret
	}
	return nil
}

func HashAndSalt(secret Secret) string {
	pwd := []byte(secret)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
