package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

var ErrNotFound = errors.New("not found")

func (a Account) GetById(id string) (entities.Account, error) {
	account, err := a.repository.GetById(id)
	if err != nil {
		return entities.Account{}, ErrNotFound
	}
	return account, nil
}
