package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Authentication) getToken(id string) (entities.Token, error) {
	a.logger.Info().
		Str("account_ID", id).
		Msg("getting token for id.")
	token, err := a.authenticationRepository.GetToken(id)
	if err != nil {
		a.logger.Error().Err(err).
			Str("account_ID", id).
			Msg("error occurred when was trying to get token for id.")
		return entities.Token{}, err
	}
	a.logger.Info().
		Str("account_ID", id).
		Msg("token was got with success for id.")
	return token, nil
}
