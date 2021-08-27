package account

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/rs/zerolog"
)

type Handler struct {
	useCase account.UseCase
	logger  zerolog.Logger
}

func NewHandler(useCase account.UseCase, logger zerolog.Logger) Handler {
	return Handler{
		useCase: useCase,
		logger:  logger,
	}
}
