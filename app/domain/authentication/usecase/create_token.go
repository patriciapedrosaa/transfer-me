package usecase

import (
	"github.com/golang-jwt/jwt"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"os"
)

func (a Authentication) CreateToken(login LoginInputs) (string, error) {
	validUser, err := a.CheckLogin(login)
	if !validUser {
		return "", err
	}

	account, err := a.accountRepository.GetByCpf(login.CPF)
	if err != nil {
		return "", err
	}

	token, err := entities.NewCreateToken(account.Name, account.AccountID)
	if err != nil {
		return "", err
	}

	err = a.authenticationRepository.CreateToken(token)
	if err != nil {
		return "", err
	}

	err = os.Setenv("ACCESS_SECRET", "dons deed crop fame blat lacy")
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
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}
