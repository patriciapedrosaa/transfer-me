package account

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

//go:generate moq -stub -out use_case_mock.go . UseCase

type CreateAccountInput struct {
	Name   string
	Secret string
	CPF    string
}

type UseCase interface {
	Create(ctx context.Context, input CreateAccountInput) (entities.Account, error)
	GetAccounts(ctx context.Context) ([]entities.Account, error)
	GetBalance(ctx context.Context, id string) (int, error)
	GetById(ctx context.Context, id string) (entities.Account, error)
	GetByCpf(ctx context.Context, cpf string) (entities.Account, error)
	UpdateBalance(ctx context.Context, originAccountId, destinationAccountId string, amount int) error
}
