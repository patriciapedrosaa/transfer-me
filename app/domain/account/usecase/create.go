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
	a.logger.Info().Msgf("Validating if CPF %s already exists...", input.CPF)
	accountExist, _ := a.repository.GetByCpf(input.CPF)
	if accountExist.CPF != "" {
		a.logger.Error().Err(ErrAlreadyExist).Msgf("CPF %s already exists on database", accountExist.CPF)
		return entities.Account{}, ErrAlreadyExist
	}

	a.logger.Info().Msgf("Creating account for CPF: %s...", input.CPF)
	newAccount, err := entities.NewCreateAccount(input.Name, vos.Secret(input.Secret), vos.CPF(input.CPF))
	if err != nil {
		a.logger.Error().Err(err).Msgf("Could not create account for CPF %s", input.CPF)
		return entities.Account{}, err
	}

	err = a.repository.CreateAccount(newAccount)
	if err != nil {
		a.logger.Error().Err(err).Msgf("Could not create account for CPF %s", input.CPF)
		return entities.Account{}, err
	}

	a.logger.Info().Msgf("Account for CPF %s was created successfully! Account ID: %s", input.CPF, newAccount.AccountID)
	return newAccount, nil
}
