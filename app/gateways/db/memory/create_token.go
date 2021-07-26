package memory

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (m *MemoryStorage) CreateToken(token entities.Token) error {
	storedToken := Token{
		ID:        token.ID,
		Name:      token.Name,
		Subject:   token.Subject,
		Issuer:    token.Issuer,
		IssuedAt:  token.IssuedAt,
		ExpiredAt: token.ExpiredAt,
	}
	m.storageAuthentication[token.ID] = storedToken
	return nil
}
