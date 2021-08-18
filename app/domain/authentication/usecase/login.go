package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

var (
	ErrInvalidCPF    = errors.New("cpf is wrong")
	ErrInvalidSecret = errors.New("secret is wrong")
)

func (a Authentication) checkLogin(inputs authentication.LoginInputs) (bool, error) {
	isValidCPF := inputs.CPF == string(inputs.Account.CPF)
	if !isValidCPF {
		return false, ErrInvalidCPF
	}
	err := vos.CompareHashAndSecret(inputs.Secret, string(inputs.Account.Secret))
	if err != nil {
		return false, ErrInvalidSecret
	}
	return true, err
}
