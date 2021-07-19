package memory

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (m MemoryStorage) CreateTransfer(transfer entities.Transfer) error {
	storedTransfer := Transfer{
		id:                   transfer.TransferID,
		accountOriginID:      transfer.AccountOriginID,
		accountDestinationID: transfer.AccountDestinationID,
		amount:               transfer.Amount,
		createdAt:            transfer.CreatedAt,
	}
	m.storageTransfer[transfer.TransferID] = storedTransfer
	return nil
}
