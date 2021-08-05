package usecase

func (a Account) UpdateBalance(originAccountId, destinationAccountId string, amount int) error {
	originAccount, err := a.repository.GetById(originAccountId)
	if err != nil {
		return err
	}
	destinationAccount, err := a.repository.GetById(destinationAccountId)
	if err != nil {
		return err
	}

	updatedOriginBalance := originAccount.Balance - amount
	updatedDestinationBalance := destinationAccount.Balance + amount

	err = a.repository.UpdateBalance(originAccountId, updatedOriginBalance)
	if err != nil {
		return err
	}
	err = a.repository.UpdateBalance(destinationAccountId, updatedDestinationBalance)
	if err != nil {
		updatedOriginBalance = updatedOriginBalance + amount
		a.repository.UpdateBalance(originAccountId, updatedOriginBalance)
		return err
	}
	return nil
}
