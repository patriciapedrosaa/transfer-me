package usecase

import (
	"context"
)

func (a Account) GetBalance(ctx context.Context, id string) (int, error) {
	a.logger.Info().
		Str("account_ID", id).
		Msg("getting balance for account id")
	account, err := a.GetById(ctx, id)
	if err != nil {
		a.logger.Error().Err(err).
			Str("account_ID", id).
			Msg("error occurred when was trying get balance for id")
		return 0, err
	}
	return account.Balance, nil
}
