package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Account) GetAccounts() ([]entities.Account, error) {
	a.logger.Info().Msg("getting accounts.")
	accounts, err := a.repository.GetAccounts()
	if err != nil {
		a.logger.Error().Err(err).Msg("error occurred when trying get accounts")
		return nil, err
	}
	return accounts, nil
}
