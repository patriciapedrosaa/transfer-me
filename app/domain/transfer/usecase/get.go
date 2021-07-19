package usecase

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

func (t Transfer) GetTransfers() ([]entities.Transfer, error) {
	transfers, err := t.transferRepository.GetTransfers()
	if err != nil {
		return nil, err
	}
	return transfers, nil
}
