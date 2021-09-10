package usecase

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Account) GetAccounts(ctx context.Context) ([]entities.Account, error) {
	a.logger.Info().Msg("getting accounts.")
	accounts, err := a.repository.GetAccounts(ctx)
	if err != nil {
		a.logger.Error().Err(err).Msg("error occurred when trying to get accounts")
		return nil, err
	}
	return accounts, nil
}
