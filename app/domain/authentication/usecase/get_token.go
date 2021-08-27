package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Authentication) getToken(id string) (entities.Token, error) {
	a.logger.Info().Msgf("Getting token for id %s", id)
	token, err := a.authenticationRepository.GetToken(id)
	if err != nil {
		a.logger.Error().Err(err).Msgf("Occurred when was trying to get token for id %s", id)
		return entities.Token{}, err
	}
	a.logger.Info().Msgf("Token was got with success for id %s", id)
	return token, nil
}
