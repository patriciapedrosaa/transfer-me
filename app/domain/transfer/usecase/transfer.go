package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
)

type Transfer struct {
	transferRepository transfer.Repository
}

func NewTransferUseCase(transferRepository transfer.Repository) Transfer {
	return Transfer{
		transferRepository: transferRepository,
	}
}
