package usecase

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

func (a Account) GetByCpf(cpf string) (entities.Account, error) {
	account, err := a.repository.GetByCpf(cpf)
	if err != nil {
		return entities.Account{}, err
	}
	return account, nil
}
