package memory

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

func (m MemoryStorage) GetTransfers() ([]entities.Transfer, error) {
	var transfers []entities.Transfer

	for _, v := range m.storageTransfer {
		var transfer entities.Transfer

		transfer.TransferID = v.id
		transfer.AccountOriginID = v.accountOriginID
		transfer.AccountDestinationID = v.accountDestinationID
		transfer.Amount = v.amount
		transfer.CreatedAt = v.createdAt

		transfers = append(transfers, transfer)
	}
	return transfers, nil
}
