package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
	"github.com/rs/zerolog"
)

type Transfer struct {
	transferRepository transfer.Repository
	logger             zerolog.Logger
}

func NewTransferUseCase(transferRepository transfer.Repository, logger zerolog.Logger) Transfer {
	return Transfer{
		transferRepository: transferRepository,
		logger:             logger,
	}
}
