package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

type CreateTransferInput struct {
	OriginAccountId      string
	DestinationAccountId string
	Amount               int
}

func (t Transfer) Create(input CreateTransferInput) (entities.Transfer, error) {
	originAccount, err := t.accountRepository.GetById(input.OriginAccountId)
	if err != nil {
		return entities.Transfer{}, err
	}
	destinationAccount, err := t.accountRepository.GetById(input.DestinationAccountId)
	if err != nil {
		return entities.Transfer{}, err
	}
	transfer, err := entities.NewCreateTransfers(originAccount, destinationAccount, input.Amount)
	if err != nil {
		return entities.Transfer{}, err
	}
	err = t.transferRepository.CreateTransfer(transfer, originAccount.AccountID)
	if err != nil {
		return entities.Transfer{}, err
	}
	return transfer, nil
}
