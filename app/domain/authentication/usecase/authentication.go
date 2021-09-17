package usecase

import (
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/rs/zerolog"
)

type Authentication struct {
	authenticationRepository authentication.Repository
	accessSecret             string
	logger                   zerolog.Logger
}

func NewAuthenticationUseCase(authenticationRepository authentication.Repository, accessSecret string, logger zerolog.Logger) Authentication {
	return Authentication{
		authenticationRepository: authenticationRepository,
		accessSecret:             accessSecret,
		logger:                   logger,
	}
}
