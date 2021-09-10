package account

import (
	"context"
)

const updateBalanceQuery = `UPDATE accounts	SET balance = $1 WHERE id = $2;`

func (r Repository) UpdateBalance(ctx context.Context, id string, value int) error {
	res, err := r.conn.Exec(ctx, updateBalanceQuery, value, id)
	if err != nil {
		return err
	}
	rowsUpdated := res.RowsAffected()
	if rowsUpdated == 0 {
		return err
	}
	return nil
}
