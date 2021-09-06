package authentication

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

//go:generate moq -stub -out use_case_mock.go . UseCase

type LoginInputs struct {
	CPF     string
	Secret  string
	Account entities.Account
}

type UseCase interface {
	CreateToken(ctx context.Context, login LoginInputs) (string, error)
	ValidatesToken(ctx context.Context, tokenString string) (entities.Token, error)
}
