package account

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
)

type Repository interface {
	CreateAccount(account entities.Account) error
	GetByCpf(cpf string) (entities.Account, error)
	GetAccounts() ([]entities.Account, error)
	UpdateBalance(cpf vos.CPF, value int) error
}
