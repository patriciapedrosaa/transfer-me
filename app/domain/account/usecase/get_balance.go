package usecase

func (a Account) GetBalance(id string) (int, error) {
	a.logger.Info().Msgf("Getting balance for account id: %s...", id)
	account, err := a.GetById(id)
	if err != nil {
		a.logger.Error().Err(err).Msgf("Occurred when was trying get balance for id %s", id)
		return 0, err
	}
	return account.Balance, nil
}
