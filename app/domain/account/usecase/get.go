package usecase

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

func (a Account) GetAccount() ([]entities.Account, error) {
	accounts, err := a.repository.GetAccounts()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
