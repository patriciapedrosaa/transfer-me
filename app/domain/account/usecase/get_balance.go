package usecase

func (a Account) GetBalance(cpf string) (int, error) {
	account, err := a.GetByCpf(cpf)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}
