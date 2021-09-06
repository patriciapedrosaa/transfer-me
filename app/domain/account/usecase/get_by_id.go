package usecase

import (
	"context"
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

var ErrNotFound = errors.New("not found")

func (a Account) GetById(ctx context.Context, id string) (entities.Account, error) {
	log := a.logger.With().Str("account_ID", id).Logger()
	log.Info().Msg("getting account by id.")
	account, err := a.repository.GetById(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg(" error occurred when was trying get account for id.")
		return entities.Account{}, ErrNotFound
	}
	log.Info().Msg("account was got with success!")
	return account, nil
}
