package authentication

import (
	"context"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

//go:generate moq -stub -out repository_mock.go . Repository

type Repository interface {
	CreateToken(ctx context.Context, token entities.Token) error
	GetToken(ctx context.Context, id string) (entities.Token, error)
}
