package memory

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (m MemoryStorage) GetByCpf(cpf string) (entities.Account, error) {
	accounts, err := m.GetAccounts()
	if err != nil {
		return entities.Account{}, err
	}
	for _, account := range accounts {
		if string(account.CPF) == cpf {
			return entities.Account{
				AccountID: account.AccountID,
				Name:      account.Name,
				CPF:       account.CPF,
				Secret:    account.Secret,
				Balance:   account.Balance,
				CreatedAt: account.CreatedAt,
			}, nil
		}
	}
	return entities.Account{}, errNotFound
}
