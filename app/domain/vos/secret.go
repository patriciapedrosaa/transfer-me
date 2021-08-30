package vos

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidSecret            = errors.New("invalid secret")
	ErrDifferentHashAndPassword = errors.New("password hashes are not same")
)

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
		return ""
	}
	return string(hash)
}

func CompareHashAndSecret(secret string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	if err != nil {
		return ErrDifferentHashAndPassword
	}
	return nil
}
