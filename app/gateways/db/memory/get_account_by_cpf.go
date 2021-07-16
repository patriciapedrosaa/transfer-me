package memory

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

var errNotFound = errors.New("not found")

func (m MemoryStorage) GetByCpf(cpf string) (entities.Account, error) {
	account := m.storageAccount[cpf]
	if account.cpf == "" {
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
