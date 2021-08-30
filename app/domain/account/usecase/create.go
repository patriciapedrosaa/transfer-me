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
	a.logger.Info().
		Str("CPF terminated in", input.CPF[len(input.CPF)-3:]).
		Msg("validating if CPF already exists.")
	accountExist, _ := a.repository.GetByCpf(input.CPF)
	if accountExist.CPF != "" {
		a.logger.Error().Err(ErrAlreadyExist).
			Str("CPF terminated in", input.CPF[len(input.CPF)-3:]).
			Msg("CPF already exists on database")
		return entities.Account{}, ErrAlreadyExist
	}

	a.logger.Info().
		Str("CPF terminated in", input.CPF[len(input.CPF)-3:]).
		Msg("creating account for CPF.")
	newAccount, err := entities.NewCreateAccount(input.Name, vos.Secret(input.Secret), vos.CPF(input.CPF))
	if err != nil {
		a.logger.Error().Err(err).
			Str("CPF terminated in", input.CPF[len(input.CPF)-3:]).
			Msgf("could not create account for CPF.")
		return entities.Account{}, err
	}

	err = a.repository.CreateAccount(newAccount)
	if err != nil {
		a.logger.Error().Err(err).
			Str("CPF terminated in", input.CPF[len(input.CPF)-3:]).
			Msg("could not create account for CPF.")
		return entities.Account{}, err
	}

	return newAccount, nil
}
