package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

var (
	ErrAlreadyExist = errors.New("account already exist")
)

func (a Account) Create(input account.CreateAccountInput) (entities.Account, error) {
	accountExist, _ := a.repository.GetByCpf(input.CPF)
	if accountExist.CPF != "" {
		return entities.Account{}, ErrAlreadyExist
	}

	newAccount, err := entities.NewCreateAccount(input.Name, vos.Secret(input.Secret), vos.CPF(input.CPF))
	if err != nil {
		return entities.Account{}, err
	}

	err = a.repository.CreateAccount(newAccount)
	if err != nil {
		return entities.Account{}, err
	}
	return newAccount, nil
}
