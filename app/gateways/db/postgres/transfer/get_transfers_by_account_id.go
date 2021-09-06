package transfer

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

const getTransferByIDQuery = `
			SELECT 
				id, 
				origin_account_id, 
				destination_account_id, 
				amount, 
				created_at
			FROM 
				transfers 
			WHERE 
				origin_account_id=$1`

func (r Repository) GetTransfersByAccountID(ctx context.Context, accountID string) ([]entities.Transfer, error) {
	rows, err := r.conn.Query(ctx, getTransferByIDQuery, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []entities.Transfer
	for rows.Next() {
		var t entities.Transfer
		if err := rows.Scan(&t.TransferID, &t.AccountOriginID, &t.AccountDestinationID, &t.Amount, &t.CreatedAt); err != nil {
			return nil, err
		}
		transfers = append(transfers, t)
	}
	return transfers, nil
}
