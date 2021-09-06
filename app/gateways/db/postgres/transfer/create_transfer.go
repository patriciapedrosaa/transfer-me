package transfer

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

const createTransferQuery = `
	INSERT INTO 
		transfers (
			id, 
			origin_account_id, 
			destination_account_id, 
			amount, 
			created_at)
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING 
		id`

func (r Repository) CreateTransfer(ctx context.Context, transfer entities.Transfer, accountID string) error {
	err := r.conn.QueryRow(ctx, createTransferQuery,
		transfer.TransferID,
		transfer.AccountOriginID,
		transfer.AccountDestinationID,
		transfer.Amount,
		transfer.CreatedAt,
	).Scan(&transfer.TransferID)

	if err != nil {
		return err
	}
	return nil
}
