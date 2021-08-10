package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

var ErrInvalidCredentials = errors.New("incorrect username or password")

func (a Authentication) CheckLogin(inputs authentication.LoginInputs) (bool, error) {
	isValidCPF := inputs.CPF == string(inputs.Account.CPF)
	if !isValidCPF {
		return false, ErrInvalidCredentials
	}
	err := vos.CompareHashAndSecret(inputs.Secret, string(inputs.Account.Secret))
	if err != nil {
		return false, ErrInvalidCredentials
	}
	return true, err
}
