package usecase

import (
	"context"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

var (
	ErrAlreadyExist = errors.New("account already exist")
)

func (a Account) Create(ctx context.Context, input account.CreateAccountInput) (entities.Account, error) {
	log := a.logger.With().Str("CPF", "********"+input.CPF[len(input.CPF)-3:]).Logger()

	log.Info().Msg("validating if CPF already exists.")
	accountExist, _ := a.repository.GetByCpf(ctx, input.CPF)
	if accountExist.CPF != "" {
		log.Error().Err(ErrAlreadyExist).Msg("CPF already exists on database")
		return entities.Account{}, ErrAlreadyExist
	}

	log.Info().Msg("creating account for CPF.")
	newAccount, err := entities.NewCreateAccount(input.Name, vos.Secret(input.Secret), vos.CPF(input.CPF))
	if err != nil {
		log.Error().Err(err).Msgf("could not create account for CPF.")
		return entities.Account{}, err
	}

	err = a.repository.CreateAccount(ctx, newAccount)
	if err != nil {
		log.Error().Err(err).Msg("could not create account for CPF.")
		return entities.Account{}, err
	}

	return newAccount, nil
}
