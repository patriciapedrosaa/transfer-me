package account

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

const createAccountQuery = `
	INSERT INTO 
		accounts (
			id, 
			name, 
			cpf, 
			secret, 
			balance, 
			created_at)
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING 
		id`

func (r Repository) CreateAccount(ctx context.Context, account entities.Account) error {
	err := r.conn.QueryRow(ctx, createAccountQuery,
		account.AccountID,
		account.Name,
		account.CPF,
		account.Secret,
		account.Balance,
		account.CreatedAt,
	).Scan(&account.AccountID)

	if err != nil {
		return err
	}
	return nil
}
