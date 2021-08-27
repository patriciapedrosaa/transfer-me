package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (t Transfer) GetTransfersByAccountID(accountID string) ([]entities.Transfer, error) {
	t.logger.Info().Msgf("Getting transfers for account id: %s...", accountID)
	transfers, err := t.transferRepository.GetTransfersByAccountID(accountID)
	if err != nil {
		t.logger.Error().Err(err).Msgf("Occurred when was trying to get transfer for account id: %s.", accountID)
		return []entities.Transfer{}, err
	}
	return transfers, nil
}
