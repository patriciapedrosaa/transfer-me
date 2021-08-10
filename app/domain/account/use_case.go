package account

import "github.com/patriciapedrosaa/transfer-me/app/domain/entities"

//go:generate moq -stub -out use_case_mock.go . UseCase

type CreateAccountInput struct {
	Name   string
	Secret string
	CPF    string
}

type UseCase interface {
	Create(input CreateAccountInput) (entities.Account, error)
	GetAccounts() ([]entities.Account, error)
	GetBalance(id string) (int, error)
	GetById(id string) (entities.Account, error)
	GetByCpf(cpf string) (entities.Account, error)
	UpdateBalance(originAccountId, destinationAccountId string, amount int) error
}
