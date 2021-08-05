package memory

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

var errNotFound = errors.New("not found")

func (m MemoryStorage) GetById(id string) (entities.Account, error) {
	account, ok := m.storageAccount[id]
	if !ok {
		return entities.Account{}, errNotFound
	}

	entityAccount := entities.Account{
		AccountID: account.id,
		Name:      account.name,
		CPF:       vos.CPF(account.cpf),
		Secret:    vos.Secret(account.secret),
		Balance:   account.balance,
		CreatedAt: account.createdAt,
	}
	return entityAccount, nil
}
