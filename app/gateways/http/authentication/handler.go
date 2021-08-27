package authentication

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/rs/zerolog"
)

type Handler struct {
	useCase        authentication.UseCase
	accountUseCase account.UseCase
	logger         zerolog.Logger
}

func NewHandler(useCase authentication.UseCase, accountUseCase account.UseCase, logger zerolog.Logger) Handler {
	return Handler{
		useCase:        useCase,
		accountUseCase: accountUseCase,
		logger:         logger,
	}
}
