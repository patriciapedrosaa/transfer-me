package account

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

type Repository interface {
	CreateAccount(account entities.Account) error
	GetById(id string) (entities.Account, error)
	GetByCpf(cpf string) (entities.Account, error)
	GetAccounts() ([]entities.Account, error)
	UpdateBalance(id string, value int) error
}
