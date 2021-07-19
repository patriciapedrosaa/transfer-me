package transfer

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

type Repository interface {
	GetTransfers() ([]entities.Transfer, error)
	CreateTransfer(transfer entities.Transfer) error
	UpdateBalance(account entities.Account, value int) error
}
