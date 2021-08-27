package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/rs/zerolog"
	"os"
)

type Authentication struct {
	authenticationRepository authentication.Repository
	accessSecret             string
	logger                   zerolog.Logger
}

func NewAuthenticationUseCase(authenticationRepository authentication.Repository, logger zerolog.Logger) Authentication {
	return Authentication{
		authenticationRepository: authenticationRepository,
		accessSecret:             os.Getenv("ACCESS_SECRET"),
		logger:                   logger,
	}
}
