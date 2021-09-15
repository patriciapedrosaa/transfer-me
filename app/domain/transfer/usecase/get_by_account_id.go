package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

var ErrInvalidId = errors.New("id format is invalid")

func (t Transfer) GetTransfersByAccountID(ctx context.Context, accountID string) ([]entities.Transfer, error) {
	log := t.logger.With().
		Str("account_ID", accountID).
		Logger()

	log.Info().Msg("getting transfers for account id")

	_, err := uuid.Parse(accountID)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to parse id")
		return []entities.Transfer{}, ErrInvalidId
	}

	transfers, err := t.transferRepository.GetTransfersByAccountID(ctx, accountID)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to get transfer for account id")
		return []entities.Transfer{}, err
	}
	return transfers, nil
}
