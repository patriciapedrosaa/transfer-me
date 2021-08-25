package transfer

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/transfer"
)

type Handler struct {
	useCase        transfer.UseCase
	accountUseCase account.UseCase
}

func NewHandler(useCase transfer.UseCase, accountUseCase account.UseCase) Handler {
	return Handler{
		useCase:        useCase,
		accountUseCase: accountUseCase,
	}
}
