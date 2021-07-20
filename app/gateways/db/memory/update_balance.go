package memory

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

func (m MemoryStorage) UpdateBalance(cpf vos.CPF, value int) error {
	_, err := m.GetByCpf(string(cpf))
	if err != nil {
		return err
	}
	accountStored := m.storageAccount[string(cpf)]
	accountStored.balance = value
	m.storageAccount[string(cpf)] = accountStored

	return nil
}
