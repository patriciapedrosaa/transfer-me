package authentication

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

const createTokenQuery = `
	INSERT INTO 
		tokens (
			id, 
			name, 
			subject, 
			issuer, 
			issued_at, 
			expired_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING 
		id`

func (r Repository) CreateToken(ctx context.Context, token entities.Token) error {
	err := r.conn.QueryRow(ctx, createTokenQuery,
		token.ID,
		token.Name,
		token.Subject,
		token.Issuer,
		token.IssuedAt,
		token.ExpiredAt,
	).Scan(&token.ID)

	if err != nil {
		return err
	}
	return nil
}
