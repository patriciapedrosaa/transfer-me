package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (t Transfer) GetTransfersByAccountID(accountID string) ([]entities.Transfer, error) {
	transfers, err := t.transferRepository.GetTransfersByAccountID(accountID)
	if err != nil {
		return []entities.Transfer{}, err
	}
	return transfers, nil
}
