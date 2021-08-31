package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Account) GetByCpf(cpf string) (entities.Account, error) {
	log := a.logger.With().Str("CPF", "********"+cpf[len(cpf)-3:]).Logger()

	log.Info().Msg("getting account by CPF.")
	account, err := a.repository.GetByCpf(cpf)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying get account for cpf.")
		return entities.Account{}, err
	}
	log.Info().Msg("account was get with success!")
	return account, nil
}
