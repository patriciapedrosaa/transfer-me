package vos

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidCPF = errors.New("invalid cpf")
	regexCpf      = regexp.MustCompile(`^(\d{11})$`)
)

type CPF string

func IsValidCPF(cpf CPF) error {
	if !regexCpf.MatchString(string(cpf)) {
		return ErrInvalidCPF
	}
	return nil
}
