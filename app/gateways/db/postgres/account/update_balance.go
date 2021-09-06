package account

import (
	"context"
)

const updateBalanceQuery = `UPDATE accounts	SET balance = $1 WHERE id = $2;`

func (r Repository) UpdateBalance(ctx context.Context, id string, value int) error {
	_, err := r.GetById(ctx, id)
	if err != nil {
		return err
	}

	_, err = r.conn.Query(ctx, updateBalanceQuery, value, id)
	if err != nil {
		return err
	}
	return nil
}
