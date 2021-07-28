package memory

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (m MemoryStorage) GetToken(id string) (entities.Token, error) {
	token, ok := m.storageAuthentication[id]
	if !ok {
		return entities.Token{}, errNotFound
	}

	entityToken := entities.Token{
		ID:        token.ID,
		Name:      token.Name,
		Subject:   token.Subject,
		Issuer:    token.Issuer,
		IssuedAt:  token.IssuedAt,
		ExpiredAt: token.ExpiredAt,
	}
	return entityToken, nil
}
