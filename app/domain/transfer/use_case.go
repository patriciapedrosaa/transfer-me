package transfer

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

//go:generate moq -stub -out use_case_mock.go . UseCase

type CreateTransferInput struct {
	OriginAccount      entities.Account
	DestinationAccount entities.Account
	Amount             int
}

type UseCase interface {
	Create(input CreateTransferInput) (entities.Transfer, error)
	GetTransfersByAccountID(accountID string) ([]entities.Transfer, error)
}
