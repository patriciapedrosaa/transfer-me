package usecase

import (
	"github.com/golang-jwt/jwt"
	"github.com/patriciapedrosaa/transfer-me/app/domain/authentication"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
)

func (a Authentication) CreateToken(login authentication.LoginInputs) (string, error) {
	validUser, err := a.CheckLogin(login)
	if !validUser {
		return "", err
	}

	token, err := entities.NewCreateToken(login.Account.Name, login.Account.AccountID)
	if err != nil {
		return "", err
	}

	err = a.authenticationRepository.CreateToken(token)
	if err != nil {
		return "", err
	}

	atClaims := jwt.MapClaims{
		"id":   token.ID,
		"name": token.Name,
		"sub":  token.Subject,
		"iss":  token.Issuer,
		"iat":  token.IssuedAt,
		"exp":  token.ExpiredAt,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessTokenString, err := accessToken.SignedString([]byte(a.accessSecret))

	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}
