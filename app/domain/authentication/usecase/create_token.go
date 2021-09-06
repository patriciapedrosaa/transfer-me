package usecase

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Authentication) CreateToken(ctx context.Context, login authentication.LoginInputs) (string, error) {
	validUser, err := a.checkLogin(login)
	if !validUser {
		return "", err
	}

	log := a.logger.With().Str("CPF", "********"+login.CPF[len(login.CPF)-3:]).Logger()
	log.Info().Msg("creating token for CPF.")
	token := entities.NewCreateToken(login.Account.Name, login.Account.AccountID)

	err = a.authenticationRepository.CreateToken(ctx, token)
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to create token for CPF.")
		return "", err
	}

	atClaims := jwt.MapClaims{
		"id":   token.ID,
		"name": token.Name,
		"sub":  token.Subject,
		"iss":  token.Issuer,
		"iat":  token.IssuedAt.Unix(),
		"exp":  token.ExpiredAt.Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessTokenString, err := accessToken.SignedString([]byte(a.accessSecret))
	if err != nil {
		log.Error().Err(err).Msg("error occurred when was trying to create token for CPF")
		return "", err
	}

	log.Info().Msg("token created with success for CPF!")
	return accessTokenString, nil
}
