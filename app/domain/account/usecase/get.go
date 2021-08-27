package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Account) GetAccounts() ([]entities.Account, error) {
	a.logger.Info().Msg("Getting accounts...")
	accounts, err := a.repository.GetAccounts()
	if err != nil {
		a.logger.Error().Err(err).Msg("Occurred when trying get accounts")
		return nil, err
	}
	return accounts, nil
}
