package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Authentication) getToken(id string) (entities.Token, error) {
	log := a.logger.With().Str("account_ID", id).Logger()
	log.Info().Msg("getting token for id.")
	token, err := a.authenticationRepository.GetToken(id)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to get token for id.")
		return entities.Token{}, err
	}
	log.Info().Msg("token was got with success for id.")
	return token, nil
}
