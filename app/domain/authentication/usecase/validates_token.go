package usecase

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"time"
)

var (
	ErrInvalidToken      = errors.New("token is invalid")
	ErrMethodInvalid     = errors.New("invalid signature method")
	ErrTokenNotFound     = errors.New("token not found")
	ErrDateFormatInvalid = errors.New("data format is invalid")
)

func (a Authentication) ValidatesToken(tokenString string) (entities.Token, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrMethodInvalid
		}
		return []byte(a.accessSecret), nil
	}
	jwtToken, err := jwt.Parse(tokenString, keyFunc)

	if err != nil {
		var jwtError *jwt.ValidationError
		if errors.As(err, &jwtError) {
			return entities.Token{}, jwtError.Inner
		}
		return entities.Token{}, ErrInvalidToken
	}

	claims := jwtToken.Claims.(jwt.MapClaims)
	iat, err := parseUnixToTime(claims["iat"])
	if err != nil {
		return entities.Token{}, err
	}
	exp, err := parseUnixToTime(claims["exp"])
	if err != nil {
		return entities.Token{}, err
	}

	token := entities.Token{
		ID:        claims["id"].(string),
		Name:      claims["name"].(string),
		Subject:   claims["sub"].(string),
		Issuer:    claims["iss"].(string),
		IssuedAt:  iat,
		ExpiredAt: exp,
	}

	_, err = a.getToken(token.ID)
	if err != nil {
		return entities.Token{}, ErrTokenNotFound
	}

	return token, nil
}

func parseUnixToTime(claim interface{}) (time.Time, error) {
	switch unixTimestamp := claim.(type) {
	case float64:
		return time.Unix(int64(unixTimestamp), 0), nil
	case json.Number:
		v, _ := unixTimestamp.Int64()
		return time.Unix(v, 0), nil
	default:
		return time.Time{}, ErrDateFormatInvalid
	}
}
