package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog/log"
)

func (a Account) GetByCpf(cpf string) (entities.Account, error) {
	a.logger.Info().Msgf("Getting account by cpf: %s...", cpf)
	account, err := a.repository.GetByCpf(cpf)
	if err != nil {
		log.Error().Err(err).Msgf("Occurred when was trying get account for cpf %s", cpf)
		return entities.Account{}, err
	}
	a.logger.Info().Msgf("Account for CPF %s was get with success", cpf)
	return account, nil
}
