package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Account) GetByCpf(cpf string) (entities.Account, error) {
	accountExist, err := a.repository.GetByCpf(cpf)
	if err != nil {
		return entities.Account{}, err
	}
	return accountExist, nil
}
