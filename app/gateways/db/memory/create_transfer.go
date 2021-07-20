package memory

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (m MemoryStorage) CreateTransfer(transfer entities.Transfer, accountID string) error {
	transfersList := m.storageTransfer[accountID]
		storedTransfer := Transfer{
		id:                   transfer.TransferID,
		accountOriginID:      transfer.AccountOriginID,
		accountDestinationID: transfer.AccountDestinationID,
		amount:               transfer.Amount,
		createdAt:            transfer.CreatedAt,
	}
  	transfersList = append(transfersList, storedTransfer)
	m.storageTransfer[accountID] = transfersList
	return nil
}
