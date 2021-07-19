package usecase

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

func (t Transfer) UpdateBalance(originAccount, destinationAccount entities.Account, amount int) error {
	_, err := t.accountRepository.GetByCpf(string(originAccount.CPF))
	if err != nil {
		return err
	}
	_, err = t.accountRepository.GetByCpf(string(destinationAccount.CPF))
	if err != nil {
		return err
	}

	updatedOriginBalance := originAccount.Balance - amount
	updatedDestinationBalance := originAccount.Balance + amount

	err = t.transferRepository.UpdateBalance(originAccount, updatedOriginBalance)
	if err != nil {
		return err
	}
	err = t.transferRepository.UpdateBalance(destinationAccount, updatedDestinationBalance)
	if err != nil {
		return err
	}
	return nil
}
