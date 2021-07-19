package memory

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

func (m MemoryStorage) UpdateBalance(account entities.Account, value int) error {
	_, err := m.GetByCpf(string(account.CPF))
	if err != nil {
		return err
	}
	accountStored, _ := m.storageAccount[string(account.CPF)]
	accountStored.balance = value
	m.storageAccount[string(account.CPF)] = accountStored

	return nil
}
