package entities

import (
	"errors"
	"github.com/google/uuid"
	"github.com/patriciapedrosaa/transfer-me/app/domain/vos"
	"time"
)

var (
	ErrInvalidName = errors.New("invalid name")
)

type Account struct {
	AccountID string     `json:"accountID"`
	Name      string     `json:"name"`
	CPF       vos.CPF    `json:"cpf"`
	Secret    vos.Secret `json:"secret"`
	Balance   int        `json:"balance"`
	CreatedAt time.Time  `json:"createdAt"`
}

func NewCreateAccount(name string, secret vos.Secret, cpf vos.CPF) (Account, error) {
	err := vos.IsValidCPF(cpf)
	if err != nil {
		return Account{}, err
	}

	err = isValidName(name)
	if err != nil {
		return Account{}, err
	}

	err = vos.IsValidSecret(secret)
	if err != nil {
		return Account{}, err
	}

	hash := vos.HashAndSalt(secret)
	hashedSecret := vos.Secret(hash)

	return Account{
		AccountID: uuid.Must(uuid.NewRandom()).String(),
		Name:      name,
		CPF:       cpf,
		Secret:    hashedSecret,
		Balance:   100,
		CreatedAt: time.Now(),
	}, nil
}

func isValidName(name string) error {
	if name == "" {
		return ErrInvalidName
	}
	return nil
}
