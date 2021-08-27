package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/rs/zerolog"
)

type Account struct {
	repository account.Repository
	logger     zerolog.Logger
}

func NewAccountUseCase(repository account.Repository, logger zerolog.Logger) Account {
	return Account{
		repository: repository,
		logger:     logger,
	}
}
