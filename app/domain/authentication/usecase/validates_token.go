package usecase

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

var ErrInvalidToken = errors.New("unauthorized")

func (a Authentication) ValidatesToken(tokenString string) (*jwt.Token, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(a.accessSecret), nil
	}
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}
