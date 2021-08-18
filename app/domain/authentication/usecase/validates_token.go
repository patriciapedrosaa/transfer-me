package usecase

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"time"
)

var (
	ErrInvalidToken = errors.New("unauthorized")
	ErrTokenExpired = errors.New("token expired")
)

func (a Authentication) ValidatesToken(tokenString string) (entities.Token, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(a.accessSecret), nil
	}
	jwtToken, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return entities.Token{}, err
	}

	claims := jwtToken.Claims.(jwt.MapClaims)
	token := entities.Token{
		ID:       jwtToken.Claims.(jwt.MapClaims)["id"].(string),
		Name:     claims["name"].(string),
		Subject:  claims["sub"].(string),
		Issuer:   claims["iss"].(string),
	}

	iat, err := stringToDate(claims["iat"].(string))
	if err != nil{
		return entities.Token{}, err
	}
	exp, err := stringToDate(claims["exp"].(string))
	if err != nil{
		return entities.Token{}, err
	}

	isValidDuration := checkDuration(iat, exp)

	if !isValidDuration {
		return entities.Token{}, ErrTokenExpired
	}

	token.IssuedAt = iat
	token.ExpiredAt = exp

	_, err = a.getToken(token.ID)
	if err != nil {
		return entities.Token{}, ErrInvalidToken
	}

	return token, nil
}

func stringToDate(str string) (time.Time, error){
	layout := "2006-01-02T15:04:05Z"
	t, err := time.Parse(layout, str)
	
	if err != nil {
		return time.Time{}, err
	}
	
	return t, nil
}

func checkDuration (timeStart time.Time, timeEnd time.Time) bool {
	duration := timeEnd.Sub(timeStart).Minutes()
	if duration != 15 {
		return false
	}
	return true
}
