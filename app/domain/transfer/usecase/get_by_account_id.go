package usecase

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (t Transfer) GetTransfersByAccountID(ctx context.Context, accountID string) ([]entities.Transfer, error) {
	log := t.logger.With().
		Str("account_ID", accountID).
		Logger()

	log.Info().Msg("getting transfers for account id")
	transfers, err := t.transferRepository.GetTransfersByAccountID(ctx, accountID)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to get transfer for account id")
		return []entities.Transfer{}, err
	}
	return transfers, nil
}
