package memory

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

func (m *MemoryStorage) CreateAccount(account entities.Account) error {
	storedAccount := Account{
		id:        account.AccountID,
		name:      account.Name,
		cpf:       string(account.CPF),
		secret:    string(account.Secret),
		balance:   account.Balance,
		createdAt: account.CreatedAt,
	}
	m.storageAccount[account.AccountID] = storedAccount
	return nil
}
