package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Account) GetById(id string) (entities.Account, error) {
	account, err := a.repository.GetById(id)
	if err != nil {
		return entities.Account{}, err
	}
	return account, nil
}
