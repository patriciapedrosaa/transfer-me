package account

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

const getByIDQuery = `
			SELECT 
			    id,
				name, 
				cpf, 
				secret, 
				balance, 
				created_at 
			FROM 
				accounts 
			WHERE 
				id=$1`

func (r Repository) GetById(ctx context.Context, id string) (entities.Account, error) {
	var account entities.Account
	err := r.conn.QueryRow(ctx, getByIDQuery, id).Scan(&account.AccountID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil {
		return entities.Account{}, err
	}
	return account, nil
}
