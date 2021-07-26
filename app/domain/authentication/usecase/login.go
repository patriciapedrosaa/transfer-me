package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

var ErrInvalidCredentials = errors.New("incorrect username or password")

type LoginInputs struct {
	CPF    string
	Secret string
}

func (a Authentication) CheckLogin(inputs LoginInputs) (bool, error) {
	account, err := a.accountRepository.GetByCpf(inputs.CPF)
	if err != nil {
		return false, ErrInvalidCredentials
	}
	err = vos.CompareHashAndSecret(inputs.Secret, string(account.Secret))
	if err != nil {
		return false, ErrInvalidCredentials
	}
	return true, err
}
