package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/patriciapedrosaa/transfer-me/app/domain/entities"
	"github.com/rs/zerolog/log"
	"time"
)

var (
	ErrInvalidToken      = errors.New("token is invalid")
	ErrMethodInvalid     = errors.New("invalid signature method")
	ErrTokenNotFound     = errors.New("token not found")
	ErrDateFormatInvalid = errors.New("data format is invalid")
)

func (a Authentication) ValidatesToken(ctx context.Context, tokenString string) (entities.Token, error) {
	a.logger.Info().Msg("validating token.")
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			a.logger.Error().Err(ErrMethodInvalid).Msg("error occurred when was trying to validate token")
			return nil, ErrMethodInvalid
		}
		return []byte(a.accessSecret), nil
	}
	jwtToken, err := jwt.Parse(tokenString, keyFunc)

	if err != nil {
		var jwtError *jwt.ValidationError
		if errors.As(err, &jwtError) {
			a.logger.Error().Err(err).Msg("error occurred when was trying to validate token")
			return entities.Token{}, jwtError.Inner
		}
		log.Error().Err(err).Msg("error occurred when was trying to validate token")
		return entities.Token{}, ErrInvalidToken
	}

	claims := jwtToken.Claims.(jwt.MapClaims)
	iat, err := parseUnixToTime(claims["iat"])
	if err != nil {
		a.logger.Error().Err(err).Msg("error occurred when was trying to validate token")
		return entities.Token{}, err
	}
	exp, err := parseUnixToTime(claims["exp"])
	if err != nil {
		a.logger.Error().Err(err).Msg("error occurred when was trying to validate token")
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

	_, err = a.getToken(ctx, token.ID)
	if err != nil {
		a.logger.Error().Err(err).Msg("error occurred when was trying to validate token")
		return entities.Token{}, ErrTokenNotFound
	}
	a.logger.Info().Msg("error occurred when was trying to validate token")
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
