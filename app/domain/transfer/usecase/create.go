package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
)

var (
	ErrUnexpected = errors.New("something went wrong")
)

func (t Transfer) Create(input transfer.CreateTransferInput) (entities.Transfer, error) {
	newTransfer, err := entities.NewTransfer(input.OriginAccount, input.DestinationAccount, input.Amount)
	if err != nil {
		return entities.Transfer{}, err
	}
	err = t.transferRepository.CreateTransfer(newTransfer, input.OriginAccount.AccountID)
	if err != nil {
		return entities.Transfer{}, ErrUnexpected
	}
	return newTransfer, nil
}
