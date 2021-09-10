package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

var ErrInvalidId = errors.New("id format is invalid")

func (a Authentication) getToken(ctx context.Context, id string) (entities.Token, error) {
	log := a.logger.With().Str("token_ID", id).Logger()
	log.Info().Msg("getting token for id.")

	_, err := uuid.Parse(id)
	if err != nil {
		log.Error().Err(err).Msg(" error occurred when was trying to get token for id")
		return entities.Token{}, ErrInvalidId
	}
	token, err := a.authenticationRepository.GetToken(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to get token for id.")
		return entities.Token{}, err
	}
	log.Info().Msg("token was got with success for id.")
	return token, nil
}
