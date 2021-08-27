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

	a.logger.Info().Msgf("Creating token for cpf: %s", login.CPF)
	token, err := entities.NewCreateToken(login.Account.Name, login.Account.AccountID)
	if err != nil {
		a.logger.Error().Err(err).Msgf("Occurred when was trying to create token for cpf %s", login.CPF)
		return "", err
	}

	err = a.authenticationRepository.CreateToken(token)
	if err != nil {
		a.logger.Error().Err(err).Msgf("Occurred when was trying to create token for cpf %s", login.CPF)
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
		a.logger.Error().Err(err).Msgf("Occurred when was trying to create token for cpf %s", login.CPF)
		return "", err
	}

	a.logger.Info().Msgf("Token created with success for cpf: %s", login.CPF)
	return accessTokenString, nil
}
