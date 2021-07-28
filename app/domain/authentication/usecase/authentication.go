package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"os"
)

type Authentication struct {
	authenticationRepository authentication.Repository
	accountRepository        account.Repository
	accessSecret             string
}

func NewAuthenticationUseCase(authenticationRepository authentication.Repository, accountRepository account.Repository) Authentication {
	return Authentication{
		authenticationRepository: authenticationRepository,
		accountRepository:        accountRepository,
		accessSecret:             os.Getenv("ACCESS_SECRET"),
	}
}
