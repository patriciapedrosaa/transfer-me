package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (t Transfer) GetTransfersByID(accountID string) ([]entities.Transfer, error) {
	transfers, err := t.transferRepository.GetTransfersByID(accountID)
	if err != nil {
		return []entities.Transfer{}, err
	}
	return transfers, nil
}
