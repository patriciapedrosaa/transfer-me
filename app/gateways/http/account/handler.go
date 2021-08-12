package account

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/account"
)

type Handler struct {
	useCase account.UseCase
}

func NewHandler(useCase account.UseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}
