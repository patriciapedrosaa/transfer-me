package authentication

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

//go:generate moq -stub -out use_case_mock.go . UseCase

type LoginInputs struct {
	CPF     string
	Secret  string
	Account entities.Account
}

type UseCase interface {
	CreateToken(login LoginInputs) (string, error)
	ValidatesToken(tokenString string) (entities.Token, error)
}
