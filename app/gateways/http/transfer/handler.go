package transfer

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
)

type Handler struct {
	useCase        transfer.UseCase
	authUseCase    authentication.UseCase
	accountUseCase account.UseCase
}

func NewHandler(useCase transfer.UseCase, authUseCase authentication.UseCase, accountUseCase account.UseCase) Handler {
	return Handler{
		useCase:        useCase,
		authUseCase:    authUseCase,
		accountUseCase: accountUseCase,
	}
}
