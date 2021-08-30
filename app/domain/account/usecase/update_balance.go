package usecase

func (a Account) UpdateBalance(originAccountId, destinationAccountId string, amount int) error {
	a.logger.Info().Msgf("updating origin account %s and destination account %s for amount = %v.", originAccountId, destinationAccountId, amount)
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
		a.logger.Error().Err(err).
			Str("origin_account_ID:", originAccountId).
			Str("destination_account_ID:", destinationAccountId).
			Int("amount: ", updatedOriginBalance).
			Msg("occurred when was trying update balance")
		return err
	}
	err = a.repository.UpdateBalance(destinationAccountId, updatedDestinationBalance)
	if err != nil {
		updatedOriginBalance = updatedOriginBalance + amount
		_ = a.repository.UpdateBalance(originAccountId, updatedOriginBalance)
		a.logger.Error().Err(err).
			Str("origin_account_ID:", originAccountId).
			Str("destination_account_ID:", destinationAccountId).
			Int("amount: ", updatedDestinationBalance).
			Msg("occurred when was trying update balance")
		return err
	}
	a.logger.Info().
		Str("origin_account_ID:", originAccountId).
		Str("destination_account_ID:", destinationAccountId).
		Int("amount: ", updatedDestinationBalance).
		Msg("update performed successfully!")
	return nil
}
