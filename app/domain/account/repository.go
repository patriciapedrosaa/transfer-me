package account

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

//go:generate moq -stub -out repository_mock.go . Repository

type Repository interface {
	CreateAccount(ctx context.Context, account entities.Account) error
	GetById(ctx context.Context, id string) (entities.Account, error)
	GetByCpf(ctx context.Context, cpf string) (entities.Account, error)
	GetAccounts(ctx context.Context) ([]entities.Account, error)
	UpdateBalance(ctx context.Context, id string, value int) error
}
