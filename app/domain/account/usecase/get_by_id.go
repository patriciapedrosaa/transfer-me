package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog/log"
)

var ErrNotFound = errors.New("not found")

func (a Account) GetById(id string) (entities.Account, error) {
	a.logger.Info().Msgf("Getting account by id: %s...", id)
	account, err := a.repository.GetById(id)
	if err != nil {
		log.Error().Err(err).Msgf("Occurred when was trying get account for id %s", id)
		return entities.Account{}, ErrNotFound
	}
	a.logger.Info().Msgf("Account for id %s was got with success", id)
	return account, nil
}
