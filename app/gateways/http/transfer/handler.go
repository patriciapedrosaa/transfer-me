package transfer

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
	"github.com/rs/zerolog"
)

type Handler struct {
	useCase        transfer.UseCase
	accountUseCase account.UseCase
	logger         zerolog.Logger
}

func NewHandler(useCase transfer.UseCase, accountUseCase account.UseCase, logger zerolog.Logger) Handler {
	return Handler{
		useCase:        useCase,
		accountUseCase: accountUseCase,
		logger:         logger,
	}
}
