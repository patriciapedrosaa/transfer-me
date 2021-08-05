package usecase

func (a Account) GetBalance(id string) (int, error) {
	account, err := a.GetById(id)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}
