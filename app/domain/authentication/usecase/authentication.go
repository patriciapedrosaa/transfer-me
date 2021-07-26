package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
)

type Authentication struct {
	authenticationRepository authentication.Repository
	accountRepository        account.Repository
}

func NewTransferUsecase(authenticationRepository authentication.Repository, accountRepository account.Repository) Authentication {
	return Authentication{
		authenticationRepository: authenticationRepository,
		accountRepository:        accountRepository,
	}
}
