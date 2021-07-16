package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
)

type Account struct {
	repository account.Repository
}

func NewAccountUsecase(repository account.Repository) Account {
	return Account{
		repository: repository,
	}
}
