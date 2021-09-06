package transfer

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

//go:generate moq -stub -out use_case_mock.go . UseCase

type CreateTransferInput struct {
	OriginAccount      entities.Account
	DestinationAccount entities.Account
	Amount             int
}

type UseCase interface {
	Create(ctx context.Context, input CreateTransferInput) (entities.Transfer, error)
	GetTransfersByAccountID(ctx context.Context, accountID string) ([]entities.Transfer, error)
}
