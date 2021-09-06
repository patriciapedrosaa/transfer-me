package authentication

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

const getTokenQuery = `
		SELECT 
			id, 
			name, 
			subject, 
			issuer, 
			issued_at, 
			expired_at 
		FROM 
			tokens
		WHERE 
				id=$1`

func (r Repository) GetToken(ctx context.Context, id string) (entities.Token, error) {
	var token entities.Token
	err := r.conn.QueryRow(ctx, getTokenQuery, id).Scan(&token.ID, &token.Name, &token.Subject, &token.Issuer, &token.IssuedAt, &token.ExpiredAt)
	if err != nil {
		return entities.Token{}, err
	}
	return token, nil
}
