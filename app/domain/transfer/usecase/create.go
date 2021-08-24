package usecase

import (
	"errors"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrUnexpected = errors.New("something went wrong")
)

func (t Transfer) Create(input transfer.CreateTransferInput) (entities.Transfer, error) {
	originAccount, err := t.accountRepository.GetById(input.OriginAccountId)
	if err != nil {
		return entities.Transfer{}, ErrNotFound
	}
	destinationAccount, err := t.accountRepository.GetById(input.DestinationAccountId)
	if err != nil {
		return entities.Transfer{}, ErrNotFound
	}
	transfer, err := entities.NewCreateTransfers(originAccount, destinationAccount, input.Amount)
	if err != nil {
		return entities.Transfer{}, err
	}
	err = t.transferRepository.CreateTransfer(transfer, originAccount.AccountID)
	if err != nil {
		return entities.Transfer{}, ErrUnexpected
	}
	return transfer, nil
}
