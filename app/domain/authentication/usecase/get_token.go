package usecase

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Authentication) getToken(ctx context.Context, id string) (entities.Token, error) {
	log := a.logger.With().Str("token_ID", id).Logger()
	log.Info().Msg("getting token for id.")
	token, err := a.authenticationRepository.GetToken(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to get token for id.")
		return entities.Token{}, err
	}
	log.Info().Msg("token was got with success for id.")
	return token, nil
}
