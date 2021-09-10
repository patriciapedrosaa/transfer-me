package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var ErrInvalidId = errors.New("id format is invalid")

func (a Account) GetBalance(ctx context.Context, id string) (int, error) {
	log := a.logger.With().Str("account_ID", id).Logger()
	log.Info().Msg("getting balance for account id")

	_, err := uuid.Parse(id)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to get balance for id")
		return 0, ErrInvalidId
	}
	account, err := a.GetById(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to get balance for id")
		return 0, err
	}
	return account.Balance, nil
}
