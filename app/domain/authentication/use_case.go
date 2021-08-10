package authentication

import (
	"github.com/golang-jwt/jwt"
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
	GetToken(id string) (entities.Token, error)
	CheckLogin(inputs LoginInputs) (bool, error)
	ValidatesToken(tokenString string) (*jwt.Token, error)
}
