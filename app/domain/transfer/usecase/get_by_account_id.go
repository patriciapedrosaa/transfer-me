package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (t Transfer) GetTransfersByAccountID(accountID string) ([]entities.Transfer, error) {
	t.logger.Info().
		Str("account_ID", accountID).
		Msg("getting transfers for account id")
	transfers, err := t.transferRepository.GetTransfersByAccountID(accountID)
	if err != nil {
		t.logger.Error().Err(err).
			Str("account_ID", accountID).
			Msg("error occurred when was trying to get transfer for account id")
		return []entities.Transfer{}, err
	}
	return transfers, nil
}
