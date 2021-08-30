package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog/log"
)

func (a Account) GetByCpf(cpf string) (entities.Account, error) {
	a.logger.Info().
		Str("CPF terminated in", cpf[len(cpf)-3:]).
		Msgf("getting account by CPF.")
	account, err := a.repository.GetByCpf(cpf)
	if err != nil {
		log.Error().Err(err).
			Str("CPF terminated in", cpf[len(cpf)-3:]).
			Msg("error occurred when was trying get account for cpf.")
		return entities.Account{}, err
	}
	a.logger.Info().
		Str("CPF terminated in", cpf[len(cpf)-3:]).
		Msg("account was get with success!")
	return account, nil
}
