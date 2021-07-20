package transfer

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

type Repository interface {
	CreateTransfer(transfer entities.Transfer, accountID string) error
	GetTransfersByID(accountID string) ([]entities.Transfer, error)
	UpdateBalance(cpf vos.CPF, value int) error
}
