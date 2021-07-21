package memory

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (m MemoryStorage) GetTransfersByAccountID(accountID string) ([]entities.Transfer, error) {
	transfersList, exists := m.storageTransfer[accountID]
	if !exists {
		return []entities.Transfer{}, errNotFound
	}

	var newTransfersList []entities.Transfer

	for _, t := range transfersList {
		var transfer entities.Transfer
		transfer.TransferID = t.id
		transfer.AccountOriginID = t.accountOriginID
		transfer.AccountDestinationID = t.accountDestinationID
		transfer.Amount = t.amount
		transfer.CreatedAt = t.createdAt

		newTransfersList = append(newTransfersList, transfer)
	}

	return newTransfersList, nil
}
