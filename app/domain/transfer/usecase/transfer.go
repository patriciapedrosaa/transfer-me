package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
)

type Transfer struct {
	transferRepository transfer.Repository
	accountRepository  account.Repository
}

func NewTransferUsecase(accountRepository account.Repository, transferRepository transfer.Repository) Transfer {
	return Transfer{
		transferRepository: transferRepository,
		accountRepository:  accountRepository,
	}
}
