package usecase

import (
	"github.com/golang-jwt/jwt"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Authentication) CreateToken(login authentication.LoginInputs) (string, error) {
	validUser, err := a.checkLogin(login)
	if !validUser {
		return "", err
	}

	a.logger.Info().
		Str("CPF terminated in", login.CPF[len(login.CPF)-3:]).
		Msg("creating token for CPF.")
	token, err := entities.NewCreateToken(login.Account.Name, login.Account.AccountID)
	if err != nil {
		a.logger.Error().Err(err).
			Str("CPF terminated in", login.CPF[len(login.CPF)-3:]).
			Msg("error occurred when was trying to create token for CPF.")
		return "", err
	}

	err = a.authenticationRepository.CreateToken(token)
	if err != nil {
		a.logger.Error().Err(err).
			Str("CPF terminated in", login.CPF[len(login.CPF)-3:]).
			Msg("error occurred when was trying to create token for CPF.")
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
		a.logger.Error().Err(err).
			Str("CPF terminated in", login.CPF[len(login.CPF)-3:]).
			Msg("error occurred when was trying to create token for CPF")
		return "", err
	}

	a.logger.Info().
		Str("CPF terminated in", login.CPF[len(login.CPF)-3:]).
		Msg("token created with success for CPF!")
	return accessTokenString, nil
}
