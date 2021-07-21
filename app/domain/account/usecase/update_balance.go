package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

func (a Account) UpdateBalance(originAccountCPF, destinationAccountCPF vos.CPF, amount int) error {
	originAccount, err := a.repository.GetByCpf(string(originAccountCPF))
	if err != nil {
		return err
	}
	destinationAccount, err := a.repository.GetByCpf(string(destinationAccountCPF))
	if err != nil {
		return err
	}

	updatedOriginBalance := originAccount.Balance - amount
	updatedDestinationBalance := destinationAccount.Balance + amount

	err = a.repository.UpdateBalance(originAccountCPF, updatedOriginBalance)
	if err != nil {
		return err
	}
	err = a.repository.UpdateBalance(destinationAccountCPF, updatedDestinationBalance)
	if err != nil {
		updatedOriginBalance = updatedOriginBalance + amount
		a.repository.UpdateBalance(originAccountCPF, updatedOriginBalance)
		return err
	}
	return nil
}
