package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

func (t Transfer) UpdateBalance(originAccountCPF, destinationAccountCPF vos.CPF, amount int) error {
	originAccount, err := t.accountRepository.GetByCpf(string(originAccountCPF))
	if err != nil {
		return err
	}
	destinationAccount, err := t.accountRepository.GetByCpf(string(destinationAccountCPF))
	if err != nil {
		return err
	}

	updatedOriginBalance := originAccount.Balance - amount
	updatedDestinationBalance := destinationAccount.Balance + amount

	err = t.transferRepository.UpdateBalance(originAccountCPF, updatedOriginBalance)
	if err != nil {
		return err
	}
	err = t.transferRepository.UpdateBalance(destinationAccountCPF, updatedDestinationBalance)
	if err != nil {
		updatedOriginBalance =updatedOriginBalance + amount
		t.transferRepository.UpdateBalance(originAccountCPF, updatedOriginBalance)
		return err
	}
	return nil
}
