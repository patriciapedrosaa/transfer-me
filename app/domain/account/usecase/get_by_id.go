package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog/log"
)

var ErrNotFound = errors.New("not found")

func (a Account) GetById(id string) (entities.Account, error) {
	a.logger.Info().
		Str("account_ID", id).
		Msg("getting account by id.")
	account, err := a.repository.GetById(id)
	if err != nil {
		log.Error().Err(err).Str("account_ID", id).Msg(" error occurred when was trying get account for id.")
		return entities.Account{}, ErrNotFound
	}
	a.logger.Info().Str("account_ID", id).Msg("account was got with success!")
	return account, nil
}
