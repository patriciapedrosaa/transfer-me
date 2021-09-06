package account

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

const getAccountByCPFQuery = `
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
				cpf=$1`

func (r Repository) GetByCpf(ctx context.Context, cpf string) (entities.Account, error) {
	var account entities.Account
	err := r.conn.QueryRow(ctx, getAccountByCPFQuery, cpf).Scan(&account.AccountID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil {
		return entities.Account{}, err
	}
	return account, nil
}
