package transfer

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

//go:generate moq -stub -out repository_mock.go . Repository

type Repository interface {
	CreateTransfer(ctx context.Context, transfer entities.Transfer, accountID string) error
	GetTransfersByAccountID(ctx context.Context, accountID string) ([]entities.Transfer, error)
}
