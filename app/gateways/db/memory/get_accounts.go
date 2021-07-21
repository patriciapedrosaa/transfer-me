package memory

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
	"sort"
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
	sort.Slice(accounts[:], func(i, j int) bool {
		return accounts[i].CPF < accounts[j].CPF
	})

	return accounts, nil
}
