package transfer

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

type Repository interface {
	CreateTransfer(transfer entities.Transfer, accountID string) error
	GetTransfersByAccountID(accountID string) ([]entities.Transfer, error)
}
