package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"os"
)

type Authentication struct {
	authenticationRepository authentication.Repository
	accessSecret             string
}

func NewAuthenticationUseCase(authenticationRepository authentication.Repository) Authentication {
	return Authentication{
		authenticationRepository: authenticationRepository,
		accessSecret:             os.Getenv("ACCESS_SECRET"),
	}
}
