package memory

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

func (m MemoryStorage) GetAccounts() ([]entities.Account, error) {
	var accounts []entities.Account

	for _, v := range m.storageAccount {
		var account entities.Account

		account.AccountID = v.id
		account.Name = v.name
		account.CPF = vos.CPF(v.cpf)
		account.Secret = vos.Secret(v.secret)
		account.Balance = v.balance
		account.CreatedAt = v.createdAt

		accounts = append(accounts, account)
	}
	return accounts, nil
}
