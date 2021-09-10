package account

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

const getAccountsQuery = `
		SELECT 
			id, 
			name, 
			cpf, 
			secret, 
			balance, 
			created_at 
		FROM 
			accounts
			`

func (r Repository) GetAccounts(ctx context.Context) ([]entities.Account, error) {
	rows, err := r.conn.Query(ctx, getAccountsQuery)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var accounts []entities.Account
	for rows.Next() {
		var a entities.Account
		if err := rows.Scan(&a.AccountID, &a.Name, &a.CPF, &a.Secret, &a.Balance, &a.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}
