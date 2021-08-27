package usecase

func (a Account) UpdateBalance(originAccountId, destinationAccountId string, amount int) error {
	a.logger.Info().Msgf("Updating origin account %s and destination account %s for amount = %v...", originAccountId, destinationAccountId, amount)
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
		a.logger.Error().Err(err).Msgf("Occurred when was trying update balance for id %s with amount = %d", originAccountId, updatedOriginBalance)
		return err
	}
	err = a.repository.UpdateBalance(destinationAccountId, updatedDestinationBalance)
	if err != nil {
		updatedOriginBalance = updatedOriginBalance + amount
		_ = a.repository.UpdateBalance(originAccountId, updatedOriginBalance)
		a.logger.Error().Err(err).Msgf("Occurred when was trying update balance for id %s with amount = %d", destinationAccountId, updatedDestinationBalance)
		return err
	}
	a.logger.Info().Msgf("Update performed successfully for origin account %s and destination account %s with amount = %v", originAccountId, destinationAccountId, amount)
	return nil
}
