package authentication

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
)

const (
	JsonContentType = "application/json"
)

type Handler struct {
	useCase        authentication.UseCase
	accountUseCase account.UseCase
}

func NewHandler(useCase authentication.UseCase, accountUseCase account.UseCase) Handler {
	return Handler{
		useCase:        useCase,
		accountUseCase: accountUseCase,
	}
}
