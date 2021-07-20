package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

type CreateTransferInput struct {
	OriginAccountCPF      string
	DestinationAccountCPF string
	Amount                int
}

func (t Transfer) Create(input CreateTransferInput) (entities.Transfer, error) {
	originAccount, err := t.accountRepository.GetByCpf(input.OriginAccountCPF)
	if err != nil {
		return entities.Transfer{}, err
	}
	destinationAccount, err := t.accountRepository.GetByCpf(input.DestinationAccountCPF)
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
	err = t.UpdateBalance(originAccount.CPF, destinationAccount.CPF, input.Amount)
	if err != nil {
		return entities.Transfer{}, err
	}
	return transfer, nil
}
